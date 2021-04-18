package numSpecial

func NumSpecial(mat [][]int) int {
	rows, cols := make([]int, len(mat)), make([]int, len(mat[0]))
	points := make([][]int, 0)
	for i, arr := range mat {
		for j, v := range arr {
			if v == 1 {
				points = append(points, []int{i, j})
				rows[i]++
				cols[j]++
			}
		}
	}
	res := 0
	for _, point := range points {
		if rows[point[0]] == 1 && cols[point[1]] == 1 {
			res++
		}
	}
	return res
}