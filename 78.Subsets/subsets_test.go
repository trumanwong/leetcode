package main

import (
	"fmt"
	"testing"
	"truman.com/leetcode/78.Subsets/subsets"
)

func TestSubsets(t *testing.T)  {
	tests := []struct{
		input []int
		output [][]int
	}{
		{[]int{1,2,3},[][]int{
			{},{1},{1,2},{1,2,3},{1,3},{2,3},{3},
		}},
		{[]int{0},[][]int{}},
	}

	for _, test := range tests {
		ret := subsets.Subsets(test.input)
		fmt.Println(ret)
	}
}