package generate

func Generate(numRows int) [][]int {
	var res [][]int
	res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			temp := 1
			if j != 0 && j != i {
				temp = res[i-1][j-1] + res[i-1][j]
			}
			res[i][j] = temp
		}
	}
	return res
}
