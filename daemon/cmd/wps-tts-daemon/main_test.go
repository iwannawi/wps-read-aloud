//go:build linux

package main

import (
	"strings"
	"testing"
)

func TestPreprocessFanchenTextSpeaksAsciiCharacters(t *testing.T) {
	got := preprocessFanchenText("深度学习是AI的核心技术，使用Python 3.11进行WPS开发。", 1.2)
	want := "深度学习是 诶 爱 的核心技术， 使用 批 歪 提 艾尺 欧 恩 三 点 一 一 进行 达不溜 批 艾丝 开发。"
	if got != want {
		t.Fatalf("preprocessed text = %q, want %q", got, want)
	}
}

func TestPreprocessFanchenTextAddsPauseOnlyForSemanticPunctuation(t *testing.T) {
	got := preprocessFanchenText("他说：“你好，WPS！”《标题》继续。", 1.2)
	want := "他说： “你好， 达不溜 批 艾丝！ ”《标题》继续。"
	if got != want {
		t.Fatalf("preprocessed text = %q, want %q", got, want)
	}
}

func TestPreprocessFanchenTextSpeaksMathSymbols(t *testing.T) {
	got := preprocessFanchenText("面积S=a*b+3/2，结果>=10%。", 1.2)
	for _, want := range []string{"艾丝", "诶", "乘", "必", "加", "三", "除", "二", "大于等于", "一 零", "百分号"} {
		if !strings.Contains(got, want) {
			t.Fatalf("preprocessed text = %q, missing %q", got, want)
		}
	}
}

func TestPreprocessFanchenTextHandlesTocPageEntries(t *testing.T) {
	cases := map[string][]string{
		"1": {"第", "一", "页"},
		"···················· 4":  {"第", "四", "页"},
		"第一章  总则\t............ 1": {"第一章", "总则", "第", "一", "页"},
	}
	for input, wants := range cases {
		got := preprocessFanchenText(input, 1.2)
		if strings.Contains(got, "·") || strings.Contains(got, "....") {
			t.Fatalf("preprocessed TOC text = %q, still contains leader dots", got)
		}
		for _, want := range wants {
			if !strings.Contains(got, want) {
				t.Fatalf("preprocessed TOC text = %q, missing %q", got, want)
			}
		}
	}
}

func TestTtsTextCandidatesAddContextForShortCJK(t *testing.T) {
	got := ttsTextCandidates("目录")
	if len(got) < 2 || got[1] != "第 目录 项" {
		t.Fatalf("ttsTextCandidates for short CJK = %#v, want contextual retry", got)
	}
}

func TestSilencePCMUsesExpectedDuration(t *testing.T) {
	wav := wavPCM{channels: 1, sampleRate: 16000, bitsPerSample: 16}
	got := silencePCM(wav, sentenceEndPauseMs(1.2))
	want := 16000 * 2 * sentenceEndPauseMsAtBaseRate / 1000
	if len(got) != want {
		t.Fatalf("silence length = %d, want %d", len(got), want)
	}
}

func TestPauseDurationsScaleWithRate(t *testing.T) {
	tests := []struct {
		rate              float64
		wantStandardMs    int
		wantSentenceEndMs int
	}{
		{rate: 1.2, wantStandardMs: 400, wantSentenceEndMs: 600},
		{rate: 0.75, wantStandardMs: 640, wantSentenceEndMs: 960},
		{rate: 1.0, wantStandardMs: 480, wantSentenceEndMs: 720},
		{rate: 1.5, wantStandardMs: 320, wantSentenceEndMs: 480},
	}
	for _, tt := range tests {
		if got := standardPauseMs(tt.rate); got != tt.wantStandardMs {
			t.Fatalf("standardPauseMs(%v) = %d, want %d", tt.rate, got, tt.wantStandardMs)
		}
		if got := sentenceEndPauseMs(tt.rate); got != tt.wantSentenceEndMs {
			t.Fatalf("sentenceEndPauseMs(%v) = %d, want %d", tt.rate, got, tt.wantSentenceEndMs)
		}
	}
}

func TestPrefetchCountUsesDynamicTextWindow(t *testing.T) {
	rs := &readSession{sentences: []ReadSentence{
		{Text: strings.Repeat("一", 101)},
		{Text: "第二句"},
		{Text: "第三句"},
	}}
	if got := rs.prefetchCount(0); got != 1 {
		t.Fatalf("prefetchCount for long first sentence = %d, want 1", got)
	}

	rs.sentences = []ReadSentence{
		{Text: strings.Repeat("一", 30)},
		{Text: strings.Repeat("二", 30)},
		{Text: strings.Repeat("三", 30)},
		{Text: strings.Repeat("四", 30)},
	}
	if got := rs.prefetchCount(0); got != 4 {
		t.Fatalf("prefetchCount for short sentences = %d, want 4", got)
	}
}
