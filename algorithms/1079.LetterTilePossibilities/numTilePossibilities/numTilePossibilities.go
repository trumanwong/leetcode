package numTilePossibilities

func NumTilePossibilities(tiles string) int {
	m := make(map[string]int)
	tileBytes := []byte(tiles)
	var recursive func(indexs []int, s string)
	recursive = func(indexs []int, s string) {
		if len(s) == len(tiles) {
			return
		}
		for i, v := range tileBytes {
			if inArray(i, indexs) {
				continue
			}
			m[s + string(v)]++
			indexs = append(indexs, i)
			recursive(indexs, s + string(v))
			indexs = indexs[:len(indexs) - 1]
		}
	}
	recursive([]int{},"")
	return len(m)
}

func inArray(needle int, stack []int) bool {
	for _, v := range stack {
		if v == needle {
			return true
		}
	}
	return false
}