package hasPathSum

import (
	"testing"
	. "truman.com/leetcode/112.PathSum/hasPathSum"
)

func TestHasPathSum(t *testing.T)  {
	root := TreeNode{5,
		&TreeNode{4,
			&TreeNode{11,&TreeNode{7,nil,nil},
				&TreeNode{2,nil,nil}},nil},
		&TreeNode{8,&TreeNode{13,nil,nil},
			&TreeNode{4,&TreeNode{5,nil,nil},&TreeNode{1,nil,nil}}}}
	tests := []struct {
		root TreeNode
		sum int
		output bool
	}{
		{root,22, true},
		{root,50,false},
	}
	for _, test := range tests {
		ret := HasPathSum(&test.root, test.sum)
		if ret != test.output {
			t.Errorf("Got %t; expected %t", ret, test.output)
		}
	}
}