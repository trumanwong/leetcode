package addStrings

import (
	"math"
	"strconv"
)

func AddStrings(num1 string, num2 string) string {
	res := ""
	diff := len(num1) - len(num2)
	diff_abs := int(math.Abs(float64(diff)))
	zeros := ""
	for i := 0; i < diff_abs; i++ {
		zeros += "0"
	}
	if diff < 0 {
		num1 = zeros + num1
	} else {
		num2 = zeros + num2
	}

	length := len(num1)
	isCarry := false
	for i := length - 1; i >= 0; i-- {
		temp1, _ := strconv.Atoi(string(num1[i]))
		temp2, _ := strconv.Atoi(string(num2[i]))
		add := temp1 + temp2
		if isCarry {
			add += 1
		}
		if add >= 10 {
			isCarry = true
			add = add % 10
		} else {
			isCarry = false
		}
		res = strconv.Itoa(add) + res
	}
	if isCarry {
		res = "1" + res
	}
	return res
}
