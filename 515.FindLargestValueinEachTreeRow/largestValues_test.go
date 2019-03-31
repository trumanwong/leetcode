package main

import (
	"testing"
	. "truman.com/leetcode/515.FindLargestValueinEachTreeRow/largestValues"
)

func TestLargestValues(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output []int
	}{
		{TreeNode{1,
			&TreeNode{3,
				&TreeNode{5,nil,nil},
				&TreeNode{3,nil,nil}},
				&TreeNode{2,nil,&TreeNode{9,nil,nil}}},
				[]int{1,3,9}},
	}
	for _, test := range tests {
		ret := LargestValues(&test.root)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}