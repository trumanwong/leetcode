package sortedArrayToBST

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := &TreeNode{nums[mid], nil, nil}
	nums = append(nums[:mid], nums[mid+1:]...)
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid:])
	return root
}
