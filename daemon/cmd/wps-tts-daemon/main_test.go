//go:build linux

package main

import "testing"

func TestPreprocessFanchenTextSpacesEnglishLetters(t *testing.T) {
	got := preprocessFanchenText("深度学习是AI的核心技术，使用Python进行开发。")
	want := "深度学习是 A I 的核心技术，使用 P y t h o n 进行开发。"
	if got != want {
		t.Fatalf("preprocessed text = %q, want %q", got, want)
	}
}
