package findJudge

func findJudge(N int, trust [][]int) int {
	count, trustOne := make([]int, N), make([]int, N)
	for _, v := range trust {
		out, in := v[0], v[1]
		count[out - 1]++
		trustOne[in - 1]++
	}

	for i := 0; i < N; i++ {
		if count[i] == 0 && trustOne[i] == N - 1 {
			return i + 1
		}
	}
	return -1
}