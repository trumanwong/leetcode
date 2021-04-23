package reverseVowels

func ReverseVowels(s string) string {
	vowel := []byte("aeiouAEIOU")
	left, right := 0, len(s)-1
	l, r := -1, -1
	S := []byte(s)
	for left < right {
		if inArray(S[left], vowel) {
			l = left
		} else {
			left++
		}
		if inArray(S[right], vowel) {
			r = right
		} else {
			right--
		}
		if l != -1 && r != -1 {
			S[l], S[r] = S[r], S[l]
			l, r = -1, -1
			left++
			right--
		}
	}
	return string(S)
}

func inArray(needle byte, hayStack []byte) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}
