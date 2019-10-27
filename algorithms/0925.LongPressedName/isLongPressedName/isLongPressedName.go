package isLongPressedName

func IsLongPressedName(name string, typed string) bool {
	n, t := []byte(name), []byte(typed)

	i, j := 0, 0
	for i < len(t) && j < len(n) {
		if t[i] != n[j] {
			n = append(n[:j], append([]byte{t[i]}, n[j:]...)...)
		}

		i++
		j++
	}

	if len(t) > len(n) {
		for i = len(n); i < len(t); i++ {
			if t[i] != n[len(n) - 1] {
				return false
			}
		}
	}

	if len(t) <= len(n) && typed != string(n) {
		return false
	}
	return true
}