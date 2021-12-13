package replaceDigits

func replaceDigits(s string) string {
	res := []byte(s)
	for i := 1; i < len(s); i += 2 {
		res[i-1] = s[i-1]
		res[i] = s[i] + s[i-1] - '0'
	}
	return string(res)
}
