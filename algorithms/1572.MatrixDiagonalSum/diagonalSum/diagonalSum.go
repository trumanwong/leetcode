package diagonalSum

func DiagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		res += mat[i][i]
		if len(mat) == 1 || (i == len(mat)-i-1) {
			continue
		}
		res += mat[i][len(mat)-i-1]
	}
	return res
}
