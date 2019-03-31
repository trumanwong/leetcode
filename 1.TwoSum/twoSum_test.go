package main

import (
	"testing"
	"truman.com/leetcode/1.TwoSum/twoSum"
)

func TestTwoSum(t *testing.T)  {
	tests := []struct{
		input []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, test := range tests {
		ret := twoSum.TwoSum(test.input, test.target)
		for i := 0; i < len(ret); i++ {
			if ret[i] != test.output[i] {
				t.Errorf("Got %v for input %v, target = %d; expected %v", ret, test.input, test.target, test.output)
			}
		}
	}
}