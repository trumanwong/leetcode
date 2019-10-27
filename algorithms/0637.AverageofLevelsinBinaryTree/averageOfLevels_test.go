package main

import (
	"leetcode/algorithms/0637.AverageofLevelsinBinaryTree/averageOfLevels"
	"testing"
)

func TestAverageOfLevels(t *testing.T)  {
	tests := []struct{
		root   averageOfLevels.TreeNode
		output []float64
	}{
		{averageOfLevels.TreeNode{3,
			&averageOfLevels.TreeNode{9,nil,nil},
			&averageOfLevels.TreeNode{20,
				&averageOfLevels.TreeNode{15,nil,nil},
				&averageOfLevels.TreeNode{7,nil,nil}}}, []float64{3,14.5,11}},
	}
	for _, test := range tests {
		ret := averageOfLevels.AverageOfLevels(&test.root)
		for i, v := range ret {
			if v - test.output[i] != 0 {
				t.Errorf("Got %v; expected %v", ret, test.output)
				break
			}
		}
	}
}