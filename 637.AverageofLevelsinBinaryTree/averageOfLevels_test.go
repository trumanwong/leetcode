package main

import (
	"testing"
	. "truman.com/leetcode/637.AverageofLevelsinBinaryTree/averageOfLevels"
)

func TestAverageOfLevels(t *testing.T)  {
	tests := []struct{
		root TreeNode
		output []float64
	}{
		{TreeNode{3,
			&TreeNode{9,nil,nil},
			&TreeNode{20,
				&TreeNode{15,nil,nil},
				&TreeNode{7,nil,nil}}},[]float64{3,14.5,11}},
	}
	for _, test := range tests {
		ret := AverageOfLevels(&test.root)
		for i, v := range ret {
			if v - test.output[i] != 0 {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}