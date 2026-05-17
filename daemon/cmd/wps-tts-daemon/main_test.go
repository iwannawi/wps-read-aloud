//go:build linux

package main

import (
	"strings"
	"testing"
)

func TestPreprocessFanchenTextSpeaksAsciiCharacters(t *testing.T) {
	got := preprocessFanchenText("深度学习是AI的核心技术，使用Python 3.11进行WPS开发。")
	want := "深度学习是 诶 爱 的核心技术， 使用 批 歪 提 艾尺 欧 恩 三 点 一 一 进行 达不溜 批 艾丝 开发。"
	if got != want {
		t.Fatalf("preprocessed text = %q, want %q", got, want)
	}
}

func TestPreprocessFanchenTextAddsPauseOnlyForSemanticPunctuation(t *testing.T) {
	got := preprocessFanchenText("他说：“你好，WPS！”《标题》继续。")
	want := "他说： “你好， 达不溜 批 艾丝！ ”《标题》继续。"
	if got != want {
		t.Fatalf("preprocessed text = %q, want %q", got, want)
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
