package convertToBase7

import "strconv"

func ConvertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res, isNegative := "", false
	if num < 0 {
		num = -num
		isNegative = true
	}

	for num != 0 {
		digit := num % 7
		num /= 7
		res = strconv.Itoa(digit) + res
	}
	if isNegative {
		res = "-" + res
	}
	return res
}
