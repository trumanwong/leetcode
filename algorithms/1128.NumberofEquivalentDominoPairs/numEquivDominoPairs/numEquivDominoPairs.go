package numEquivDominoPairs

import "strconv"

func NumEquivDominoPairs(dominoes [][]int) int {
	m, res := make(map[string]int), 0
	for _, v := range dominoes {
		str1 := strconv.Itoa(v[0] * 10 + v[1])
		str2 := strconv.Itoa(v[1] * 10 + v[0])
		if m[str1] == 0 && m[str2] == 0 {
			m[str1]++
		} else if m[str1] != 0 {
			m[str1]++
		} else {
			m[str2]++
		}
	}
	for _, v := range m {
		for i := 1; i <= v - 1; i++ {
			res +=i
		}
	}
	return res
}