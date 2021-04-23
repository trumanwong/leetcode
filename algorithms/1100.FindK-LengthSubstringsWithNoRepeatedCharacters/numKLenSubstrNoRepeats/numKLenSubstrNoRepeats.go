package numKLenSubstrNoRepeats

func NumKLenSubstrNoRepeats(S string, K int) int {
	start, ret := 0, 0
	s := []byte(S)
	for start < len(S)-K+1 {
		end := start + K
		m := make(map[byte]int)
		ret++
		for i := start; i < end; i++ {
			m[s[i]]++
			if m[s[i]] > 1 {
				ret--
				break
			}
		}
		start++
	}
	return ret
}
