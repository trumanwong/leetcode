package main

import (
	"leetcode/algorithms/0665.Non-decreasingArray/checkPossibility"
	"testing"
)

func TestCheckPossibility(t *testing.T)  {
	tests := []struct{
		nums []int
		output bool
	}{
		{[]int{4,2,3}, true},
		{[]int{4,2,1}, false},
	}

	for _, test := range tests {
		ret := checkPossibility.CheckPossibility(test.nums)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.nums, test.output)
		}
	}
}