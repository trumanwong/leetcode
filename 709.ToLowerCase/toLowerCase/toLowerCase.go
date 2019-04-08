package toLowerCase

func ToLowerCase(str string) string {
	res := ""

	for _, v := range str {
		if v >= 65 && v <= 90 {
			res += string(v + 32)
		} else {
			res += string(v)
		}
	}
	return res
}