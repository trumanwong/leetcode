package main

import (
	"leetcode/algorithms/0437.PathSumIII/pathSum"
	. "leetcode/common/treeNode"
	"testing"
)

func TestPathSum(t *testing.T)  {
	tests := []struct{
		root []interface{}
		sum int
		output int
	}{
		{[]interface{}{10,5,-3,3,2,nil,11,3,-2,nil,1}, 8, 3},
	}

	for _, test := range tests {
		ret := pathSum.PathSum(Constructor(0, test.root), test.sum)
		if ret != test.output {
			t.Errorf("Got %d for input %v, sum = %d; expected %d", ret, test.root, test.sum, test.output)
		}
	}
}