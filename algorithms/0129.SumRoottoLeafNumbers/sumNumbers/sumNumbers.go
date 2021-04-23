package sumNumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumNumbers(root *TreeNode) int {
	res := 0
	var recursiveNumbers func(*TreeNode, int)
	recursiveNumbers = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += sum
			return
		}
		recursiveNumbers(node.Left, sum)
		recursiveNumbers(node.Right, sum)
	}
	recursiveNumbers(root, 0)
	return res
}
