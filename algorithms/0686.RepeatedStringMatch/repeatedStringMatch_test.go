package main

import (
	"leetcode/algorithms/0686.RepeatedStringMatch/repeatedStringMatch"
	"testing"
)

func TestRepeatedStringMatch(t *testing.T)  {
	tests := []struct{
		A string
		B string
		output int
	}{
		{"abcd", "cdabcdab", 3},
	}

	for _, test := range tests {
		ret := repeatedStringMatch.RepeatedStringMatch(test.A, test.B)
		if ret != test.output {
			t.Errorf("Got %d for input %s, %s; expected %d", ret, test.A, test.B, test.output)
		}
	}
}