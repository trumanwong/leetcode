package main

import (
	"testing"
	"truman.com/leetcode/167.TwoSumII-Inputarrayissorted/twoSum"
)

func TestTwoSum(t *testing.T)  {
	tests := []struct{
		numbers []int
		target int
		output []int
	}{
		{[]int{2,7,11,15},9,[]int{1,2}},
	}

	for _, test := range tests {
		ret := twoSum.TwoSum(test.numbers, test.target)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input numbers = %v, target = %d; expected %v", ret, test.numbers, test.target, test.output)
			}
		}
	}
}