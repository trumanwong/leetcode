package maxLengthBetweenEqualCharacters

import "math"

func MaxLengthBetweenEqualCharacters(s string) int {
	m := make(map[int32]int)
	res := -1
	for i, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = i
		} else {
			res = int(math.Max(float64(res), float64(i)-float64(m[v])-1))
		}
	}
	return res
}
