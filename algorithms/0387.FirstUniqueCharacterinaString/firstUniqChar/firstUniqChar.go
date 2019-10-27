package firstUniqChar

func FirstUniqChar(s string) int {
	m := make(map[int32]int)
	for _, v := range s {
		m[v]++
	}
	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}