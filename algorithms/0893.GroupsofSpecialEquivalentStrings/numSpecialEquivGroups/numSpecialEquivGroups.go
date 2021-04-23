package numSpecialEquivGroups

func NumSpecialEquivGroups(A []string) int {
	seen := make(map[string]int)
	for _, s := range A {
		count := make([]rune, 52)
		for i, v := range s {
			count[int(v-'a')+26*(i%2)]++
		}
		seen[string(count)]++
	}
	return len(seen)
}
