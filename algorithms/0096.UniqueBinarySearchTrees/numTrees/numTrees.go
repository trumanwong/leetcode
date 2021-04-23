package numTrees

func NumTrees(n int) int {
	if n < 3 {
		return n
	}
	res := make([]int, n+1)
	res[0], res[1], res[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		sum := 0
		for j := 1; j <= i; j++ {
			sum += res[j-1] * res[i-j]
		}
		res[i] = sum
	}
	return res[n]
}
