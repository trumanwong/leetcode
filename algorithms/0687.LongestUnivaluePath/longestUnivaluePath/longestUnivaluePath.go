package longestUnivaluePath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func LongestUnivaluePath(root *TreeNode) int {
	ans = 0
	arrowLength(root)
	return ans
}

func arrowLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := arrowLength(root.Left)
	right := arrowLength(root.Right)
	arrowLeft, arrowRight := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		arrowLeft += left + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		arrowRight += right + 1
	}
	ans = max(ans, arrowLeft+arrowRight)
	return max(arrowLeft, arrowRight)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
