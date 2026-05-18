package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Config struct {
	Listen string
	Sherpa SherpaConfig
}

type SherpaConfig struct {
	Bin              string
	NumThreads       int
	TargetSampleRate int
	VitsModel        string
	VitsLexicon      string
	VitsTokens       string
	VitsRuleFsts     string
	VitsSpeakerID    int
}

type ReadSentence struct {
	Text string `json:"text"`
}

type ReadStartRequest struct {
	Sentences []ReadSentence `json:"sentences"`
	Rate      float64        `json:"rate"`
	Prefetch  int            `json:"prefetch"`
}

type Server struct {
	root    string
	cfg     Config
	mu      sync.Mutex
	session *Session
	current *exec.Cmd
}

type Session struct {
	id      string
	server  *Server
	ctx     context.Context
	cancel  context.CancelFunc
	mu      sync.Mutex
	state   string
	message string
	index   int
	total   int
}

const (
	pauseBaseRate                = 1.2
	standardPauseMsAtBaseRate   = 400
	sentenceEndPauseMsAtBaseRate = 600
)

var asciiTokenRE = regexp.MustCompile(`[A-Za-z0-9]+(?:[._+\-][A-Za-z0-9]+)*`)

func main() {
	configPath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	root, _ := os.Getwd()
	cfg := defaultConfig()
	if err := loadSimpleYAML(*configPath, &cfg); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Printf("load config failed, using defaults: %v", err)
	}
	cfg = absolutizeConfig(root, cfg)
	server := &Server{root: root, cfg: cfg}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", server.health)
	mux.HandleFunc("/selftest", server.selftest)
	mux.HandleFunc("/read/start", server.readStart)
	mux.HandleFunc("/read/status", server.readStatus)
	mux.HandleFunc("/read/stop", server.stop)
	mux.HandleFunc("/stop", server.stop)
	mux.HandleFunc("/audio/probe", server.audioProbe)
	mux.HandleFunc("/docs/", server.docs)
	mux.HandleFunc("/", server.web)
	log.Printf("wps-tts-daemon-windows listening on http://%s", cfg.Listen)
	if err := http.ListenAndServe(cfg.Listen, cors(mux)); err != nil {
		log.Fatal(err)
	}
}

func defaultConfig() Config {
	return Config{
		Listen: "127.0.0.1:19860",
		Sherpa: SherpaConfig{
			Bin:              "engines/sherpa-onnx/sherpa-onnx-offline-tts.exe",
			NumThreads:       2,
			TargetSampleRate: 16000,
			VitsModel:        "voices/sherpa/vits-zh-hf-fanchen-C/vits-zh-hf-fanchen-C.onnx",
			VitsLexicon:      "voices/sherpa/vits-zh-hf-fanchen-C/lexicon.txt",
			VitsTokens:       "voices/sherpa/vits-zh-hf-fanchen-C/tokens.txt",
			VitsRuleFsts:     "voices/sherpa/vits-zh-hf-fanchen-C/phone.fst,voices/sherpa/vits-zh-hf-fanchen-C/date.fst,voices/sherpa/vits-zh-hf-fanchen-C/number.fst,voices/sherpa/vits-zh-hf-fanchen-C/new_heteronym.fst",
			VitsSpeakerID:    14,
		},
	}
}

func absolutizeConfig(root string, cfg Config) Config {
	cfg.Sherpa.Bin = abs(root, cfg.Sherpa.Bin)
	cfg.Sherpa.VitsModel = abs(root, cfg.Sherpa.VitsModel)
	cfg.Sherpa.VitsLexicon = abs(root, cfg.Sherpa.VitsLexicon)
	cfg.Sherpa.VitsTokens = abs(root, cfg.Sherpa.VitsTokens)
	var fsts []string
	for _, item := range strings.Split(cfg.Sherpa.VitsRuleFsts, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			fsts = append(fsts, abs(root, item))
		}
	}
	cfg.Sherpa.VitsRuleFsts = strings.Join(fsts, ",")
	return cfg
}

func abs(root, path string) string {
	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}
	return filepath.Join(root, filepath.FromSlash(path))
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	engine := "none"
	if fileExists(s.cfg.Sherpa.Bin) && fileExists(s.cfg.Sherpa.VitsModel) && fileExists(s.cfg.Sherpa.VitsLexicon) && fileExists(s.cfg.Sherpa.VitsTokens) {
		engine = "sherpa-onnx"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":           engine == "sherpa-onnx",
		"version":      appVersion(s.root),
		"engine":       engine,
		"audio_player": "Windows SoundPlayer",
		"message":      "Windows 本地离线朗读服务已启动。",
	})
}

func (s *Server) audioProbe(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"version":  appVersion(s.root),
		"selected": "Windows SoundPlayer",
		"results": []map[string]string{{"name": "Windows SoundPlayer", "status": "ok"}},
	})
}

func (s *Server) selftest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()
	wav, err := s.synthesize(ctx, "测试", 1.2)
	if wav != "" {
		defer os.Remove(wav)
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, friendlyError(err))
		return
	}
	info, err := os.Stat(wav)
	if err != nil || info.Size() == 0 {
		writeError(w, http.StatusInternalServerError, "语音引擎未生成有效音频。")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "bytes": info.Size()})
}

func (s *Server) readStart(w http.ResponseWriter, r *http.Request) {
	var req ReadStartRequest
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	if err := json.NewDecoder(io.LimitReader(r.Body, 8<<20)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "请求格式不正确，请重新打开 WPS 后再试。")
		return
	}
	req.Rate = clampRate(req.Rate)
	var sentences []ReadSentence
	for _, sentence := range req.Sentences {
		text := cleanText(sentence.Text)
		if text != "" {
			sentences = append(sentences, ReadSentence{Text: text})
		}
	}
	if len(sentences) == 0 {
		writeError(w, http.StatusBadRequest, "没有可朗读的有效句子。")
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	session := &Session{id: newID(), server: s, ctx: ctx, cancel: cancel, state: "preparing", message: "朗读服务正在启动，请耐心等待", index: -1, total: len(sentences)}
	s.mu.Lock()
	s.stopLocked()
	s.session = session
	s.mu.Unlock()
	go session.run(sentences, req.Rate)
	writeJSON(w, http.StatusOK, session.status())
}

func (s *Server) readStatus(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	session := s.session
	s.mu.Unlock()
	if session == nil {
		writeJSON(w, http.StatusOK, map[string]any{"state": "idle", "current_index": -1, "total": 0})
		return
	}
	writeJSON(w, http.StatusOK, session.status())
}

func (s *Server) stop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	s.mu.Lock()
	s.stopLocked()
	s.session = nil
	s.mu.Unlock()
	writeJSON(w, http.StatusOK, map[string]any{"status": "stopped"})
}

func (s *Server) web(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/addin")
	if path == "" || path == "/" {
		path = "/index.html"
	}
	http.ServeFile(w, r, filepath.Join(s.root, "addin", filepath.FromSlash(path)))
}

func (s *Server) docs(w http.ResponseWriter, r *http.Request) {
	name := filepath.Base(r.URL.Path)
	switch name {
	case "THIRD_PARTY_NOTICES.md":
		http.ServeFile(w, r, filepath.Join(s.root, "third_party_licenses", name))
	case "RELEASE_NOTES.md", "SOURCE_OFFER.md", "ACCEPTANCE_TEST.md":
		http.ServeFile(w, r, filepath.Join(s.root, name))
	default:
		http.NotFound(w, r)
	}
}

func (s *Server) stopLocked() {
	if s.session != nil {
		s.session.cancel()
	}
	if s.current != nil && s.current.Process != nil {
		_ = s.current.Process.Kill()
	}
	s.current = nil
}

func (s *Server) setCurrent(cmd *exec.Cmd) {
	s.mu.Lock()
	s.current = cmd
	s.mu.Unlock()
}

func (s *Server) clearCurrent(cmd *exec.Cmd) {
	s.mu.Lock()
	if s.current == cmd {
		s.current = nil
	}
	s.mu.Unlock()
}

func (ss *Session) run(sentences []ReadSentence, rate float64) {
	for i, sentence := range sentences {
		if ss.ctx.Err() != nil {
			ss.setState("stopped", "朗读已停止", i)
			return
		}
		ss.setState("synthesizing", "正在准备第 "+strconv.Itoa(i+1)+" 句", i)
		wav, err := ss.server.synthesize(ss.ctx, sentence.Text, rate)
		if err != nil {
			if ss.ctx.Err() != nil {
				ss.setState("stopped", "朗读已停止", i)
			} else {
				ss.setState("error", friendlyError(err), i)
			}
			return
		}
		ss.setState("playing", "正在朗读第 "+strconv.Itoa(i+1)+" 句", i)
		err = ss.server.play(ss.ctx, wav)
		_ = os.Remove(wav)
		if err != nil {
			if ss.ctx.Err() != nil {
				ss.setState("stopped", "朗读已停止", i)
			} else {
				ss.setState("error", friendlyError(err), i)
			}
			return
		}
	}
	ss.setState("done", "朗读完成", len(sentences)-1)
	ss.server.mu.Lock()
	if ss.server.session == ss {
		ss.server.session = nil
	}
	ss.server.mu.Unlock()
}

func (ss *Session) setState(state, message string, index int) {
	ss.mu.Lock()
	ss.state = state
	ss.message = message
	ss.index = index
	ss.mu.Unlock()
}

func (ss *Session) status() map[string]any {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return map[string]any{"id": ss.id, "state": ss.state, "message": ss.message, "current_index": ss.index, "total": ss.total}
}

func (s *Server) synthesize(ctx context.Context, text string, rate float64) (string, error) {
	if !fileExists(s.cfg.Sherpa.Bin) {
		return "", errors.New("no available tts engine")
	}
	tmp, err := os.CreateTemp("", "wps-read-aloud-*.wav")
	if err != nil {
		return "", err
	}
	tmpPath := tmp.Name()
	tmp.Close()
	args := []string{
		"--vits-model=" + s.cfg.Sherpa.VitsModel,
		"--vits-lexicon=" + s.cfg.Sherpa.VitsLexicon,
		"--tokens=" + s.cfg.Sherpa.VitsTokens,
		"--sid=" + strconv.Itoa(s.cfg.Sherpa.VitsSpeakerID),
		"--num-threads=" + strconv.Itoa(s.cfg.Sherpa.NumThreads),
		"--vits-noise-scale=0.667",
		"--vits-noise-scale-w=0.8",
		"--vits-length-scale=" + fmt.Sprintf("%.3f", 1/clampRate(rate)),
		"--output-filename=" + tmpPath,
	}
	if s.cfg.Sherpa.TargetSampleRate > 0 {
		args = append(args, "--tts-sample-rate="+strconv.Itoa(s.cfg.Sherpa.TargetSampleRate))
	}
	if fsts := existingRuleFsts(s.cfg.Sherpa.VitsRuleFsts); fsts != "" {
		args = append(args, "--tts-rule-fsts="+fsts)
	}
	args = append(args, preprocessFanchenText(text, rate))
	cmd := exec.CommandContext(ctx, s.cfg.Sherpa.Bin, args...)
	cmd.Dir = s.root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	s.setCurrent(cmd)
	err = cmd.Run()
	s.clearCurrent(cmd)
	if err != nil {
		os.Remove(tmpPath)
		return "", fmt.Errorf("sherpa-onnx failed: %w", err)
	}
	return tmpPath, nil
}

func (s *Server) play(ctx context.Context, wav string) error {
	script := fmt.Sprintf("(New-Object Media.SoundPlayer %q).PlaySync()", wav)
	cmd := exec.CommandContext(ctx, "powershell.exe", "-NoProfile", "-Command", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	s.setCurrent(cmd)
	err := cmd.Run()
	s.clearCurrent(cmd)
	return err
}

func preprocessFanchenText(text string, rate float64) string {
	text = strings.NewReplacer("　", " ", "℃", "摄氏度", "&", "和").Replace(text)
	text = asciiTokenRE.ReplaceAllStringFunc(text, func(token string) string {
		var parts []string
		for _, r := range token {
			if spoken := asciiCharSpeech(r); spoken != "" {
				parts = append(parts, spoken)
			}
		}
		if len(parts) == 0 {
			return " "
		}
		return " " + strings.Join(parts, " ") + " "
	})
	return normalizeTtsPunctuationSpacing(text, rate)
}

func normalizeTtsPunctuationSpacing(text string, rate float64) string {
	var out []rune
	for _, r := range text {
		switch {
		case isSemanticPausePunctuation(r):
			for len(out) > 0 && unicode.IsSpace(out[len(out)-1]) {
				out = out[:len(out)-1]
			}
			out = append(out, r, ' ')
		case unicode.IsSpace(r):
			if len(out) > 0 && out[len(out)-1] != ' ' {
				out = append(out, ' ')
			}
		default:
			out = append(out, r)
		}
	}
	return strings.TrimSpace(string(out))
}

func isSemanticPausePunctuation(r rune) bool {
	switch r {
	case '，', ',', '、', '。', '；', ';', '：', ':', '！', '!', '？', '?':
		return true
	default:
		return false
	}
}

func asciiCharSpeech(r rune) string {
	switch r {
	case '0':
		return "零"
	case '1':
		return "一"
	case '2':
		return "二"
	case '3':
		return "三"
	case '4':
		return "四"
	case '5':
		return "五"
	case '6':
		return "六"
	case '7':
		return "七"
	case '8':
		return "八"
	case '9':
		return "九"
	case '.':
		return "点"
	case '-':
		return "杠"
	case '_':
		return "下划线"
	}
	switch unicode.ToUpper(r) {
	case 'A':
		return "诶"
	case 'B':
		return "必"
	case 'C':
		return "西"
	case 'D':
		return "迪"
	case 'E':
		return "伊"
	case 'F':
		return "艾弗"
	case 'G':
		return "吉"
	case 'H':
		return "艾尺"
	case 'I':
		return "爱"
	case 'J':
		return "杰"
	case 'K':
		return "开"
	case 'L':
		return "艾勒"
	case 'M':
		return "艾姆"
	case 'N':
		return "恩"
	case 'O':
		return "欧"
	case 'P':
		return "批"
	case 'Q':
		return "丘"
	case 'R':
		return "阿尔"
	case 'S':
		return "艾丝"
	case 'T':
		return "提"
	case 'U':
		return "优"
	case 'V':
		return "维"
	case 'W':
		return "达不溜"
	case 'X':
		return "艾克斯"
	case 'Y':
		return "歪"
	case 'Z':
		return "兹"
	}
	return ""
}

func existingRuleFsts(value string) string {
	var kept []string
	for _, item := range strings.Split(value, ",") {
		item = strings.TrimSpace(item)
		if fileExists(item) {
			kept = append(kept, item)
		}
	}
	return strings.Join(kept, ",")
}

func appVersion(root string) string {
	data, err := os.ReadFile(filepath.Join(root, "version.json"))
	if err != nil {
		return "dev"
	}
	var info struct{ Version string `json:"version"` }
	if json.Unmarshal(data, &info) != nil || strings.TrimSpace(info.Version) == "" {
		return "dev"
	}
	return info.Version
}

func friendlyError(err error) string {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "no available tts engine"):
		return "朗读引擎不可用，请重新安装加载项安装包。"
	case strings.Contains(msg, "sherpa-onnx failed"):
		return "Sherpa-onnx 语音引擎启动失败，请检查安装包完整性。"
	default:
		return "朗读失败，请稍后重试。"
	}
}

func cleanText(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	var lines []string
	for _, line := range strings.Split(text, "\n") {
		lines = append(lines, strings.TrimSpace(line))
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func clampRate(rate float64) float64 {
	if rate <= 0 {
		return 1.2
	}
	if rate < 0.5 {
		return 0.5
	}
	if rate > 2.0 {
		return 2.0
	}
	return rate
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func newID() string {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return hex.EncodeToString(buf)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]any{"error": message})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loadSimpleYAML(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var section string
	for _, raw := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasSuffix(line, ":") {
			section = strings.TrimSuffix(line, ":")
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), "\"'")
		switch section + "." + key {
		case ".listen":
			cfg.Listen = value
		case "sherpa.bin":
			cfg.Sherpa.Bin = value
		case "sherpa.num_threads":
			if parsed, err := strconv.Atoi(value); err == nil {
				cfg.Sherpa.NumThreads = parsed
			}
		case "sherpa.target_sample_rate":
			if parsed, err := strconv.Atoi(value); err == nil {
				cfg.Sherpa.TargetSampleRate = parsed
			}
		case "sherpa.vits_model":
			cfg.Sherpa.VitsModel = value
		case "sherpa.vits_lexicon":
			cfg.Sherpa.VitsLexicon = value
		case "sherpa.vits_tokens":
			cfg.Sherpa.VitsTokens = value
		case "sherpa.vits_rule_fsts":
			cfg.Sherpa.VitsRuleFsts = value
		case "sherpa.vits_speaker_id":
			if parsed, err := strconv.Atoi(value); err == nil {
				cfg.Sherpa.VitsSpeakerID = parsed
			}
		}
	}
	return nil
}
