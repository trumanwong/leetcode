package main

import (
	"leetcode/algorithms/0653.TwoSumIVInputisaBST/findTarget"
	. "leetcode/common/treeNode"
	"testing"
)

func TestFindTarget(t *testing.T)  {
	tests := []struct{
		root []interface{}
		k int
		output bool
	}{
		{[]interface{}{5,3,6,2,4,nil,7}, 9, true},
		{[]interface{}{5,3,6,2,4,nil,7}, 28, false},
	}

	for _, test := range tests {
		ret := findTarget.FindTarget(Constructor(0, test.root), test.k)
		if ret != test.output {
			t.Errorf("Got %t for input %v, k = %d; expected %t", ret, test.root, test.k, test.output)
		}
	}
}