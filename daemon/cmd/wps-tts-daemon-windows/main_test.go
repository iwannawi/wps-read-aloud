package main

import (
	"strings"
	"testing"
)

func TestWindowsPreprocessSpeaksAsciiAndMathSymbols(t *testing.T) {
	got := preprocessFanchenText("面积S=a*b+3/2，结果>=10%。", 1.2)
	for _, want := range []string{"艾丝", "诶", "乘", "必", "加", "三", "除", "二", "大于等于", "一 零", "百分号"} {
		if !strings.Contains(got, want) {
			t.Fatalf("preprocessed text = %q, missing %q", got, want)
		}
	}
}

func TestWindowsPrefetchCountUsesDynamicTextWindow(t *testing.T) {
	ss := &Session{sentences: []ReadSentence{
		{Text: strings.Repeat("一", 101)},
		{Text: "第二句"},
		{Text: "第三句"},
	}}
	if got := ss.prefetchCount(0); got != 1 {
		t.Fatalf("prefetchCount for long first sentence = %d, want 1", got)
	}

	ss.sentences = []ReadSentence{
		{Text: strings.Repeat("一", 30)},
		{Text: strings.Repeat("二", 30)},
		{Text: strings.Repeat("三", 30)},
		{Text: strings.Repeat("四", 30)},
	}
	if got := ss.prefetchCount(0); got != 4 {
		t.Fatalf("prefetchCount for short sentences = %d, want 4", got)
	}
}

func TestWindowsSherpaNumThreadsIsCapped(t *testing.T) {
	cases := map[int]int{
		0: 1,
		1: 1,
		4: 4,
		8: 4,
	}
	for input, want := range cases {
		if got := sherpaNumThreads(input); got != want {
			t.Fatalf("sherpaNumThreads(%d) = %d, want %d", input, got, want)
		}
	}
}
