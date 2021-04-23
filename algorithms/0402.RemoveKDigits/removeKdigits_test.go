package main

import (
	"leetcode/algorithms/0402.RemoveKDigits/removeKdigits"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		nums   string
		k      int
		output string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
	}

	for _, test := range tests {
		ret := removeKdigits.RemoveKdigits(test.nums, test.k)
		if ret != test.output {
			t.Errorf("Got %s for nums = %s, k = %d; expected %s", ret, test.nums, test.k, test.output)
		}
	}
}
