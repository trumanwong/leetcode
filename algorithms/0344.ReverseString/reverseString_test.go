package main

import (
	"leetcode/algorithms/0344.ReverseString/reverseString"
	"testing"
)

func TestReverseString(t *testing.T)  {
	tests := []struct{
		s string
		output string
	} {
		{"hello", "olleh"},
		{"Hannah", "hannaH"},
	}

	for _, test := range tests {
		temp := []byte(test.s)
		reverseString.ReverseString(temp)
		ret := string(temp)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
