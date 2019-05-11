package main

import (
	"testing"
	"truman.com/leetcode/1035.UncrossedLines/maxUncrossedLines"
)

func TestMaxUncrossedLines(t *testing.T)  {
	tests := []struct{
		A []int
		B []int
		output int
	}{
		{[]int{1,4,2}, []int{1,2,4}, 2},
	}

	for _, test := range tests {
		ret := maxUncrossedLines.MaxUncrossedLines(test.A, test.B)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %v); expected %d", ret, test.A, test.B, test.output)
		}
	}
}