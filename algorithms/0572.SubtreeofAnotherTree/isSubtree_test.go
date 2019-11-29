package main

import (
	"leetcode/algorithms/0572.SubtreeofAnotherTree/isSubtree"
	"leetcode/common/treeNode"
	"testing"
)

func TestIsSubtree(t *testing.T)  {
	tests := []struct{
		s []interface{}
		t []interface{}
		output bool
	}{
		{[]interface{}{3, 4, 5, 1, 2}, []interface{}{4, 1, 2}, true},
	}

	for _, test := range tests {
		ret := isSubtree.IsSubtree(treeNode.Constructor(0, test.s), treeNode.Constructor(0, test.t))
		if ret != test.output {
			t.Errorf("Got %t for input %v, %v; expected %t", ret, test.s, test.t, test.output)
		}
	}
}