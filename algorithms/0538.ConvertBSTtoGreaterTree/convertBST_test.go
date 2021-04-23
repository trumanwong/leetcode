package main

import (
	"leetcode/algorithms/0538.ConvertBSTtoGreaterTree/convertBST"
	"leetcode/common/treeNode"
	"reflect"
	"testing"
)

func TestConvertBST(t *testing.T) {
	tests := []struct {
		root   []interface{}
		output []interface{}
	}{
		{[]interface{}{5, 2, 13}, []interface{}{18, 20, 13}},
	}

	for _, test := range tests {
		ret := convertBST.ConvertBST(treeNode.Constructor(0, test.root))
		if !reflect.DeepEqual(ret, treeNode.Constructor(0, test.output)) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.root, test.output)
		}
	}
}
