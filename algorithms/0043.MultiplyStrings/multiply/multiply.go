package multiply

import "strconv"

func Multiply(num1 string, num2 string) string {
	allLen := len(num1) + len(num2)
	tmpResult := make([]int, allLen)
	result := ""
	//模拟手算从最后一位开始处理
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			temp1, _ := strconv.Atoi(string(num1[i]))
			temp2, _ := strconv.Atoi(string(num2[j]))
			tmpResult[i+j+1] += temp1 * temp2
		}
	}

	//进位
	for i := allLen - 1; i > 0; i-- {
		if tmpResult[i] > 9 {
			tmpResult[i-1] += tmpResult[i] / 10
			tmpResult[i] %= 10
		}
	}

	index := -1
	//转换为字符串
	for i := 0; i <= allLen-1; i++ {
		result += strconv.Itoa(tmpResult[i])
		if tmpResult[i] == 0 || index != -1 {
			continue
		}
		if tmpResult[i] != 0 {
			index = i
		}
	}
	if index == -1 {
		return "0"
	}
	return result[index:]
}
