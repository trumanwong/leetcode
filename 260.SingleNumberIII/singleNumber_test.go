package main

import (
	"testing"
	"truman.com/leetcode/260.SingleNumberIII/singleNumberIII"
)

func TestSingleNumber(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{1,2,1,3,2,5}, []int{5,3}},
	}

	for _, test := range tests {
		ret := singleNumberIII.SingleNumber(test.input)
		for i := 0; i < len(ret); i++ {
			if ret[i] != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}