package main

import (
	"leetcode/algorithms/0844.BackspaceStringCompare/backspaceCompare"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		S      string
		T      string
		output bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a##c", "#a#c", true},
		{"a#c", "b", false},
	}

	for _, test := range tests {
		ret := backspaceCompare.BackspaceCompare(test.S, test.T)
		if ret != test.output {
			t.Errorf("Got %t for S = %s, T = %s; expected %t", ret, test.S, test.T, test.output)
		}
	}
}
