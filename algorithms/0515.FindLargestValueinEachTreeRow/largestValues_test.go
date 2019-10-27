package main

import (
	"leetcode/algorithms/0515.FindLargestValueinEachTreeRow/largestValues"
	"testing"
)

func TestLargestValues(t *testing.T)  {
	tests := []struct{
		root   largestValues.TreeNode
		output []int
	}{
		{largestValues.TreeNode{1,
			&largestValues.TreeNode{3,
				&largestValues.TreeNode{5,nil,nil},
				&largestValues.TreeNode{3,nil,nil}},
				&largestValues.TreeNode{2,nil,&largestValues.TreeNode{9,nil,nil}}},
				[]int{1,3,9}},
	}
	for _, test := range tests {
		ret := largestValues.LargestValues(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}