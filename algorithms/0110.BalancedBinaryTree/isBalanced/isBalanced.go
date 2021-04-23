package isBalanced

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	// O(n)
	//if root == nil {
	//	return true
	//}
	//balance := true
	//height(root, &balance)
	//return balance

	// O(nlogn)
	if root == nil {
		return true
	}
	left_height := height2(root.Left)
	right_height := height2(root.Right)
	balance := true
	if left_height-right_height > 1 || left_height-right_height < -1 {
		balance = false
	}
	return balance && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func height2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left_height := height2(root.Left)
	right_height := height2(root.Right)
	if left_height > right_height {
		return left_height + 1
	}
	return right_height + 1
}

func height(root *TreeNode, balance *bool) int {
	if root == nil {
		return 0
	}
	left_height := height(root.Left, balance)
	right_height := height(root.Right, balance)
	if left_height-right_height > 1 || left_height-right_height < -1 {
		*balance = false
		return -1
	}
	if left_height > right_height {
		return left_height + 1
	}
	return right_height + 1
}
