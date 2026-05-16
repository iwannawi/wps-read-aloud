//go:build linux

package main

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"
)

func TestPreparePlaybackWavAppliesVolumeOnCopy(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "source.wav")
	spec := wavPCM{format: 1, channels: 1, sampleRate: 16000, bitsPerSample: 16}
	pcm := make([]byte, 4)
	putPCM16(pcm[0:2], 8000)
	putPCM16(pcm[2:4], -8000)
	f, err := os.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := writePCM16Wav(f, spec, pcm); err != nil {
		f.Close()
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	playPath, cleanup, err := preparePlaybackWav(src, 40)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	if playPath == src {
		t.Fatal("expected volume adjustment to use a copy")
	}

	original, err := readPCM16Wav(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := int16(binary.LittleEndian.Uint16(original.data[0:2])); got != 8000 {
		t.Fatalf("source wav was modified, got first sample %d", got)
	}

	adjusted, err := readPCM16Wav(playPath)
	if err != nil {
		t.Fatal(err)
	}
	if got := int16(binary.LittleEndian.Uint16(adjusted.data[0:2])); got != 4000 {
		t.Fatalf("first sample = %d, want 4000", got)
	}
	if got := int16(binary.LittleEndian.Uint16(adjusted.data[2:4])); got != -4000 {
		t.Fatalf("second sample = %d, want -4000", got)
	}
}

func putPCM16(buf []byte, sample int16) {
	binary.LittleEndian.PutUint16(buf, uint16(sample))
}

func TestPreparePlaybackWavKeepsDefaultVolumePath(t *testing.T) {
	playPath, cleanup, err := preparePlaybackWav("/tmp/example.wav", 80)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	if playPath != "/tmp/example.wav" {
		t.Fatalf("play path = %q, want original path", playPath)
	}
}

func TestPreprocessFanchenTextSpacesEnglishLetters(t *testing.T) {
	got := preprocessFanchenText("深度学习是AI的核心技术，使用Python进行开发。")
	want := "深度学习是 A I 的核心技术，使用 P y t h o n 进行开发。"
	if got != want {
		t.Fatalf("preprocessed text = %q, want %q", got, want)
	}
}
