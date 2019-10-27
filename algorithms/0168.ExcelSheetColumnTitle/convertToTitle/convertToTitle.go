package convertToTitle

func ConvertToTitle(n int) string {
	res := ""
	for n != 0 {
		temp := n % 26 + 64
		if n % 26 == 0 {
			temp = 26 + 64
			n--
		}
		res = string(byte(temp)) + res
		n /= 26
	}
	return res
}