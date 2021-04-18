package main

import (
	"leetcode/algorithms/1609.EvenOddTree/isEvenOddTree"
	. "leetcode/common/treeNode"
	"testing"
)

func TestIsEvenOddTree(t *testing.T)  {
	tests := []struct{
		root *TreeNode
		output bool
	}{
		{Constructor(0, []interface{}{1,10,4,3,nil,7,9,12,8,6,nil,nil,2}), true},
		{Constructor(0, []interface{}{5,9,1,3,5,7}), false},
	}

	for _, test := range tests {
		ret := isEvenOddTree.IsEvenOddTree(test.root)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.root, test.output)
		}
	}
}