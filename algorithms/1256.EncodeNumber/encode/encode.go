package encode

func Encode(num int) string {
	if num == 0 {
		return ""
	}
	res := ""
	for num != 0 {
		num--
		if num%2 != 0 {
			res = "1" + res
		} else {
			res = "0" + res
		}
		num /= 2
	}
	return res
}
