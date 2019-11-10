package main

import (
	"leetcode/algorithms/1250.CheckIfItIsaGoodArray/isGoodArray"
	"testing"
)

func TestIsGoodArray(t *testing.T)  {
	tests := []struct{
		nums []int
		output bool
	}{
		{[]int{12, 5, 7, 23}, true},
		{[]int{29, 6, 10}, true},
		{[]int{3, 6}, false},
	}

	for _, test := range tests {
		ret := isGoodArray.IsGoodArray(test.nums)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.nums, test.output)
		}
	}
}