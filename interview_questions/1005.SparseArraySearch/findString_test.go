package main

import (
	"leetcode/interview_questions/1005.SparseArraySearch/findString"
	"testing"
)

func TestFindString(t *testing.T) {
	tests := []struct {
		words  []string
		s      string
		output int
	}{
		{[]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ta", -1},
		{[]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ball", 4},
	}

	for _, test := range tests {
		ret := findString.FindString(test.words, test.s)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %s); expected %d", ret, test.words, test.s, test.output)
		}
	}
}
