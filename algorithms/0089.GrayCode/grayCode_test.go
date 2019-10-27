package main

import (
	"leetcode/algorithms/0089.GrayCode/grayCode"
	"testing"
)

func TestGrayCode(t *testing.T)  {
	tests := []struct{
		n int
		output []int
	}{
		{2, []int{0, 1, 3, 2}},
		{0, []int{0}},
	}

	for _, test := range tests {
		ret := grayCode.GrayCode(test.n)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %d; expected %v", ret, test.n, test.output)
				break
			}
		}
	}
}