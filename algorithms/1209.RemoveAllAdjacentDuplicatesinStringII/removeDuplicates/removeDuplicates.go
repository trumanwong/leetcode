package removeDuplicates

func RemoveDuplicates(s string, k int) string {
	b := []byte(s)

	last := ""
	for {
		if last == string(b) {
			break
		}
		last = string(b)
		end := len(b) - k
		for i := 0; i <= end; i++ {
			l, r := i, i+k-1
			canSub := true
			for l < r {
				if b[l] != b[i] || b[r] != b[i] {
					canSub = false
					break
				}
				l++
				r--
			}

			if canSub {
				b = append(b[0:i], b[i+k:]...)
				break
			}
		}
	}
	return string(b)
}
