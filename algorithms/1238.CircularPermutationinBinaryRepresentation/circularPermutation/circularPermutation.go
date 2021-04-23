package circularPermutation

func CircularPermutation(n int, start int) []int {
	res := []int{start}
	for i := 0; i < n; i++ {
		for j := len(res) - 1; ^j != 0; j-- {
			res = append(res, int(uint(res[j])^1<<uint(i)))
		}
	}
	return res
}
