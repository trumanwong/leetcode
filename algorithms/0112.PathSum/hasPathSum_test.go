package hasPathSum

import (
	"leetcode/algorithms/0112.PathSum/hasPathSum"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	root := hasPathSum.TreeNode{5,
		&hasPathSum.TreeNode{4,
			&hasPathSum.TreeNode{11, &hasPathSum.TreeNode{7, nil, nil},
				&hasPathSum.TreeNode{2, nil, nil}}, nil},
		&hasPathSum.TreeNode{8, &hasPathSum.TreeNode{13, nil, nil},
			&hasPathSum.TreeNode{4, &hasPathSum.TreeNode{5, nil, nil}, &hasPathSum.TreeNode{1, nil, nil}}}}
	tests := []struct {
		root   hasPathSum.TreeNode
		sum    int
		output bool
	}{
		{root, 22, true},
		{root, 50, false},
	}
	for _, test := range tests {
		ret := hasPathSum.HasPathSum(&test.root, test.sum)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}
