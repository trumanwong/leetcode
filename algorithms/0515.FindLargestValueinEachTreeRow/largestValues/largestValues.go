package largestValues

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LargestValues(root *TreeNode) []int {
	res := []int{}
	var recursiveValues func(*TreeNode, int)
	recursiveValues = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) < level {
			res = append(res, node.Val)
		} else if res[level-1] < node.Val {
			res[level-1] = node.Val
		}
		level++
		recursiveValues(node.Left, level)
		recursiveValues(node.Right, level)
	}
	recursiveValues(root, 1)
	return res
}
