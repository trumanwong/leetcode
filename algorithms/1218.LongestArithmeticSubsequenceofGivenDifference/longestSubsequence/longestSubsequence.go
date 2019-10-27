package longestSubsequence

func LongestSubsequence(arr []int, difference int) int {
	m := make(map[int]int)
	res :=1;
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = m[arr[i] - difference] + 1
		res = max(res, m[arr[i]])
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}