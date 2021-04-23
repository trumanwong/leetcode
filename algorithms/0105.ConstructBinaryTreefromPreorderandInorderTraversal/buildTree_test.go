package main

import (
	"leetcode/algorithms/0105.ConstructBinaryTreefromPreorderandInorderTraversal/buildTree"
	"leetcode/common/treeNode"
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		output   []interface{}
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []interface{}{3, 9, 20, nil, nil, 15, 7}},
	}

	for _, test := range tests {
		ret := buildTree.BuildTree(test.preorder, test.inorder)
		if !reflect.DeepEqual(ret, treeNode.Constructor(0, test.output)) {
			t.Errorf("Got %v for preorder = %v, inorder = %v; expected %v", ret, test.preorder, test.inorder, test.output)
		}
	}
}
