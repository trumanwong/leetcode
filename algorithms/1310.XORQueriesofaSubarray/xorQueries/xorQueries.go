package xorQueries

func xorQueries(arr []int, queries [][]int) []int {
	dp := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		dp[i+1] = dp[i] ^ arr[i]
	}
	res := make([]int, len(queries))
	for i, v := range queries {
		res[i] = dp[v[0]] ^ dp[v[1]+1]
	}
	return res
}
