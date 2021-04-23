package monotoneIncreasingDigits

import "strconv"

func MonotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	str := []byte(strconv.Itoa(N))
	flag := len(str)
	for i := len(str) - 1; i >= 1; i-- {
		if str[i] < str[i-1] {
			flag = i
			str[i-1]--
		}
	}

	for i := flag; i < len(str); i++ {
		str[i] = '9'
	}
	ret, _ := strconv.Atoi(string(str))
	return ret
}
