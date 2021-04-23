package thousandSeparator

import "strconv"

func ThousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	i := 0
	for n != 0 {
		if i == 3 {
			res = "." + res
			i = 0
		}
		temp := n % 10
		n /= 10
		res = strconv.Itoa(temp) + res
		i++
	}
	return res
}
