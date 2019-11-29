package isSubtree

import . "leetcode/common/treeNode"

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	return isSameTree(s, t) || IsSubtree(s.Left, t) || IsSubtree(s.Right, t)
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	return s != nil && t != nil && s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}