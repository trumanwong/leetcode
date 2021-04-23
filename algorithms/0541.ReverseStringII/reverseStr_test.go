package main

import (
	"leetcode/algorithms/0541.ReverseStringII/reverseStr"
	"testing"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		output string
	}{
		{"abcdefg", 2, "bacdfeg"},
	}

	for _, test := range tests {
		ret := reverseStr.ReverseStr(test.s, test.k)
		if ret != test.output {
			t.Errorf("Got %s for input s = %s, k = %d; expected %s", ret, test.s, test.k, test.output)
		}
	}
}
