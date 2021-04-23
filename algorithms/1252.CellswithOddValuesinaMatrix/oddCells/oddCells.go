package oddCells

func OddCells(n int, m int, indices [][]int) int {
	arr := [][]int{}
	for i := 0; i < n; i++ {
		temp := make([]int, m)
		arr = append(arr, temp)
	}
	for _, indice := range indices {
		for i := 0; i < m; i++ {
			arr[indice[0]][i]++
		}
		for j := 0; j < n; j++ {
			arr[j][indice[1]]++
		}
	}
	res := 0
	for _, varr := range arr {
		for _, v := range varr {
			if v%2 != 0 {
				res++
			}
		}
	}
	return res
}
