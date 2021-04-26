package findString

func FindString(words []string, s string) int {
	l, r, mid := 0, len(words)-1, 0
	for l <= r {
		mid = (l + r) / 2
		for mid > l && words[mid] == "" {
			mid--
		}
		if words[mid] == s {
			return mid
		} else if words[mid] > s {
			r = mid - 1
		} else {
			l = (l + r) / 2 + 1
		}
	}
	return -1
}
