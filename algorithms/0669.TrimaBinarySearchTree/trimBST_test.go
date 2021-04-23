package main

import (
	"leetcode/algorithms/0669.TrimaBinarySearchTree/trimBST"
	. "leetcode/common/treeNode"
	"reflect"
	"testing"
)

func TestTrimBST(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		L      int
		R      int
		output *TreeNode
	}{
		{Constructor(0, []interface{}{1, 0, 2}), 1, 2, Constructor(0, []interface{}{1, nil, 2})},
	}

	for _, test := range tests {
		ret := trimBST.TrimBST(test.root, test.L, test.R)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %v, L = %d, R = %d; expected %v", ret, test.root, test.L, test.R, test.output)
		}
	}
}
