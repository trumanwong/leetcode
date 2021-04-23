package generateMatrix

func GenerateMatrix(n int) [][]int {
	var arr [][]int
	for i := 0; i < n; i++ {
		temp := make([]int, n)
		arr = append(arr, temp)
	}

	temp := n * n
	c, j := 1, 0
	for c <= temp {
		for i := j; i < n-j; i++ {
			arr[j][i] = c
			c++
		}

		for i := j + 1; i < n-j; i++ {
			arr[i][n-j-1] = c
			c++
		}

		for i := n - j - 2; i >= j; i-- {
			arr[n-j-1][i] = c
			c++
		}

		for i := n - j - 2; i > j; i-- {
			arr[i][j] = c
			c++
		}
		j++
	}
	return arr
}
