package largeGroupPositions

func LargeGroupPositions(S string) [][]int {
	s := []byte(S)
	ret := [][]int{}
	flag, count, start := s[0], 1, 0
	for i := 1; i < len(S); i++ {
		if flag != s[i] {
			if count >= 3 {
				ret = append(ret, []int{start, i - 1})
			}
			flag, count, start = s[i], 1, i
		} else {
			count++
		}
	}
	if count >= 3 {
		ret = append(ret, []int{start, len(s) - 1})
	}
	return ret
}