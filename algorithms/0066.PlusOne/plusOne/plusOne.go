package plusOne

func PlusOne(digits []int) []int {
	digitsSize := len(digits)
	result := make([]int, digitsSize)
	var i int
	for i = 0; i < digitsSize; i++ {
		//判断是否进位
		if digits[i] != 9 {
			break
		}
	}

	if i == digitsSize {
		//判断是否要加长输出数组长度（输入数组全为9）
		result[0] = 1
		for i = 1; i < digitsSize+1; i++ {
			if i == digitsSize {
				result = append(result, 0)
			} else {
				result[i] = 0
			}
		}
		return result
	}

	i = digitsSize - 1
	result[i] = digits[i] + 1
	for ; i > 0; i-- {
		if result[i] == 10 {
			// 判断不是第一位为9的进位
			result[i] = 0
			result[i-1] = digits[i-1] + 1
		} else {
			result[i-1] = digits[i-1]
		}
	}
	return result
}
