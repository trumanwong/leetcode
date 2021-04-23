package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var curr []int
	var recursiveSum func(*TreeNode, int)
	recursiveSum = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if node.Val == s {
				curr = append(curr, node.Val)
				temp := make([]int, len(curr))
				copy(temp, curr)
				res = append(res, temp)
				curr = curr[:len(curr)-1]
			}
			return
		}
		curr = append(curr, node.Val)
		new_sum := s - node.Val
		recursiveSum(node.Left, new_sum)
		recursiveSum(node.Right, new_sum)
		curr = curr[:len(curr)-1]
	}
	recursiveSum(root, sum)
	return res
}
