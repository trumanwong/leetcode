package toHex

func ToHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex, res := "0123456789abcdef", ""
	for num != 0 && len(res) < 8 {
		res = string(hex[num&0xf]) + res
		num >>= 4
	}
	return res
}
