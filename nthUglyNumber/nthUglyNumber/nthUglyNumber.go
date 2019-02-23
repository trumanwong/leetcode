package nthUglyNumber

func NthUglyNumber(n int) int {
	if n <= 0 || n > 1690 {
		return 0
	}
	if n == 1 {
		return 1
	}
	index2, index3, index5 := 0, 0, 0
	ret := make([]int, n)
	ret[0] = 1
	for i := 1; i < n; i++ {
		if ret[index3] * 3 < ret[index5] * 5 {
			ret[i] = ret[index3] * 3
		} else {
			ret[i] = ret[index5] * 5
		}
		if ret[i] > ret[index2] * 2 {
			ret[i] = ret[index2] * 2
		}
		if ret[i] == ret[index2] * 2 {
			index2++
		}
		if ret[i] == ret[index3] * 3 {
			index3++
		}
		if ret[i] == ret[index5] * 5 {
			index5++
		}
	}
	return ret[n - 1]
}