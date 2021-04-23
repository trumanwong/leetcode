package repeatedNTimes

func RepeatedNTimes(A []int) int {
	maps := make(map[int]int)
	res := 0
	for _, v := range A {
		if _, ok := maps[v]; ok {
			res = v
			break
		}
		maps[v] = v
	}
	return res
}
