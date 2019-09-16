package calculateTime

func calculateTime(keyboard string, word string) int {
	m, ret := make(map[rune]int), 0

	for i, v := range keyboard {
		m[v] = i
	}

	last := 0
	for _, v := range word {
		ret += abs(last - m[v])
		last = m[v]
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}