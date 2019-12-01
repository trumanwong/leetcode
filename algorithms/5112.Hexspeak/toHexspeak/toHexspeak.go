package toHexspeak

import (
	"strconv"
)

func ToHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	hex := strconv.FormatInt(int64(n), 16)
	res := ""
	for _, v := range hex {
		if v == '0' {
			res += "O"
		} else if v == '1' {
			res += "I"
		} else if v >= 'a' && v <= 'f' {
			res += string(v - 32)
		} else {
			return "ERROR"
		}
	}
	return res
}