package main

import (
	"strings"
	"testing"
)

func TestWindowsPreprocessSpeaksAsciiAndMathSymbols(t *testing.T) {
	got := preprocessFanchenText("面积S=a*b+3/2，结果>=10%，误差±0.5%，WPS Office。", 1.2)
	for _, want := range []string{"艾丝", "诶", "乘", "必", "加", "三", "除", "二", "大于等于", "百分之十", "正负", "百分之零点五", "达不溜屁挨思", "凹斐思"} {
		if !strings.Contains(got, want) {
			t.Fatalf("preprocessed text = %q, missing %q", got, want)
		}
	}
}

func TestWindowsPreprocessHandlesTocPageEntries(t *testing.T) {
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

func TestWindowsTtsTextCandidatesAddContextForShortCJK(t *testing.T) {
	got := ttsTextCandidates("目录")
	if len(got) < 2 || got[1] != "第 目录 项" {
		t.Fatalf("ttsTextCandidates for short CJK = %#v, want contextual retry", got)
	}
}

func TestWindowsCleanTextSkipsEmptyTableCellMarkers(t *testing.T) {
	if got := cleanText("\r\a"); got != "" {
		t.Fatalf("cleanText for empty table cell = %q, want empty", got)
	}
	if got := ttsTextCandidates(cleanText("\r\a")); len(got) != 0 {
		t.Fatalf("ttsTextCandidates for empty table cell = %#v, want no candidates", got)
	}
}

func TestWindowsPrefetchCountUsesDynamicTextWindow(t *testing.T) {
	ss := &Session{sentences: []ReadSentence{
		{Text: strings.Repeat("一", 241)},
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
		{Text: strings.Repeat("五", 30)},
		{Text: strings.Repeat("六", 30)},
		{Text: strings.Repeat("七", 30)},
		{Text: strings.Repeat("八", 30)},
		{Text: strings.Repeat("九", 30)},
	}
	if got := ss.prefetchCount(0); got != 8 {
		t.Fatalf("prefetchCount for short sentences = %d, want 8", got)
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
