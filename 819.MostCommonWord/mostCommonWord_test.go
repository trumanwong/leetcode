package main

import (
	"testing"
	"truman.com/leetcode/819.MostCommonWord/mostCommonWord"
)

func TestMostCommonWord(t *testing.T)  {
	tests := []struct{
		paragraph string
		banned []string
		output string
	}{
		{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"},"ball"},
	}

	for _, test := range tests {
		ret := mostCommonWord.MostCommonWord(test.paragraph, test.banned)
		if ret != test.output {
			t.Errorf("Got %s for input %s, banned %v; expected %s", ret, test.paragraph, test.banned, test.output)
		}
	}
}