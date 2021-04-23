package main

import (
	"leetcode/algorithms/1544.MakeTheStringGreat/makeGood"
	"testing"
)

func TestMakeGood(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"leEeetcode", "leetcode"},
		{"abBAcC", ""},
	}

	for _, test := range tests {
		ret := makeGood.MakeGood(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}
