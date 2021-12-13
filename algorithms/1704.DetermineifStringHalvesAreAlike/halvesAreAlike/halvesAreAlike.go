package halvesAreAlike

func halvesAreAlike(s string) bool {
	m := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	l, r := 0, len(s)-1
	countL, countR := 0, 0
	for l <= r {
		if m[s[l]] {
			countL++
		}
		if m[s[r]] {
			countR++
		}
		l++
		r--
	}
	return countL == countR
}
