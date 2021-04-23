package findLUSlength

func FindLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
