package main

import (
	"leetcode/algorithms/0400.NthDigit/findNthDigit"
	"testing"
)

func TestFindNthDigit(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{3, 3},
		{11, 0},
	}

	for _, test := range tests {
		ret := findNthDigit.FindNthDigit(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}
