package main

import (
	"leetcode/algorithms/0543.DiameterofBinaryTree/diameterOfBinaryTree"
	. "leetcode/common/treeNode"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T)  {
	tests := []struct{
		root []interface{}
		output int
	}{
		{[]interface{}{1,2,3,4,5}, 3},
	}

	for _, test := range tests {
		ret := diameterOfBinaryTree.DiameterOfBinaryTree(Constructor(0, test.root))
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.root, test.output)
		}
	}
}