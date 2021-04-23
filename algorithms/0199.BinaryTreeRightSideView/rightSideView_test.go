package main

import (
	"leetcode/algorithms/0199.BinaryTreeRightSideView/rightSideView"
	"leetcode/common/treeNode"
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		root   []interface{}
		output []int
	}{
		{[]interface{}{1, 2, 3, nil, 5, nil, 4}, []int{1, 3, 4}},
	}

	for _, test := range tests {
		ret := rightSideView.RightSideView(treeNode.Constructor(0, test.root))
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.root, test.output)
		}
	}
}
