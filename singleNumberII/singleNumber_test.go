package main

import (
	"testing"
	"truman.com/leetcode/singleNumberII/singleNumberII"
)

func TestSingleNumber(t *testing.T)  {
	tests := []struct{
		input []int
		output int
	}{
		{[]int{2,2,3,2}, 3},
		{[]int{0,1,0,1,0,1,99}, 99},
	}

	for _, test := range tests {
		ret := singleNumberII.SingleNumber(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}