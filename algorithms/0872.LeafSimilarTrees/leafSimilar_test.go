package main

import (
	"leetcode/algorithms/0872.LeafSimilarTrees/leafSimilar"
	. "leetcode/common/treeNode"
	"testing"
)

func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		root1  []interface{}
		root2  []interface{}
		output bool
	}{
		{[]interface{}{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}, []interface{}{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}, true},
	}

	for _, test := range tests {
		ret := leafSimilar.LeafSimilar(Constructor(0, test.root1), Constructor(0, test.root2))
		if ret != test.output {
			t.Errorf("Got %t for input root1=%v, root2 = %v; expected %t", ret, test.root1, test.root2, test.output)
		}
	}
}
