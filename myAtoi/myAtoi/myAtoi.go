package myAtoi

import (
	"math"
	"strconv"
	"strings"
)

func MyAtoi(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}
	_, e := strconv.Atoi(string(str[0]))
	if str == "" || len(str) == 0 || (str[0] != '-' && str[0] != '+' && e != nil) {
		return 0
	}
	if len(str) == 1 && (str[0] == '-' || str[0] == '+') {
		return 0
	}

	result := ""
	var isTop = false
	for i, ch := range str {
		temp, e := strconv.Atoi(string(ch))
		if i > 0 && e != nil {
			break;
		}
		if e == nil {
			//获取首字符
			if !isTop && temp > 0 {
				result = result + strconv.Itoa(temp)
				isTop = true
			} else if isTop {
				result = result + strconv.Itoa(temp)
			}
		} else {
			result = result + string(str[0])
		}
	}
	if result == "" || result == "-" || result == "+" {
		return 0
	}
	if len(result) > 12 {
		result = result[0:12]
	}
	ret, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	if ret > math.MaxInt32  {
		return math.MaxInt32
	}
	if ret < math.MinInt32 {
		return math.MinInt32
	}
	return ret
}
