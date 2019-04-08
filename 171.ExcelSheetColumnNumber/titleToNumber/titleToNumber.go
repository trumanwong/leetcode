package titleToNumber

func TitleToNumber(s string) int {
	res := 0
	for _, v := range s {
		res = res * 26 + int(v) - 64
	}
	return res
}