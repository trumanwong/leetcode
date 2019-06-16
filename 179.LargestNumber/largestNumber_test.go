package main

import (
	"testing"
	"truman.com/leetcode/179.LargestNumber/largestNumber"
)

func TestLargestNumber(t *testing.T)  {
	tests := []struct{
		nums []int
		output string
	}{
		{[]int{10,2},"210"},
		{[]int{3,30,34,5,9}, "9534330"},
	}

	for _, test := range tests {
		ret := largestNumber.LargestNumber(test.nums)
		if ret != test.output {
			t.Errorf("Got %s for input %v; expected %s", ret, test.nums, test.output)
		}
	}
}