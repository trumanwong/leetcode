package main

import (
	"fmt"
	"leetcode/algorithms/0090.SubsetsII/subsetsWithDup"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		input  []int
		output [][]int
	}{
		{[]int{1, 2, 2}, [][]int{
			{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2},
		}},
	}
	for _, test := range tests {
		ret := subsetsWithDup.SubsetsWithDup(test.input)
		fmt.Println(ret)
	}
}
