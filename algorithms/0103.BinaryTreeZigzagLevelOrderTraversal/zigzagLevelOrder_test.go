package main

import (
	"leetcode/algorithms/0103.BinaryTreeZigzagLevelOrderTraversal/zigzagLevelOrder"
	"leetcode/common/treeNode"
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		root   []interface{}
		output [][]int
	}{
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, [][]int{{3}, {20, 9}, {15, 7}}},
	}

	for _, test := range tests {
		ret := zigzagLevelOrder.ZigzagLevelOrder(treeNode.Constructor(0, test.root))
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.root, test.output)
		}
	}
}
