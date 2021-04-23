package maxSatisfied

func MaxSatisfied(customers []int, grumpy []int, X int) int {
	/*
	   首先 找到不改变的时候客人就满意的数量和 同时更新数组
	   这样问题就转变为 求数组指定长度最大和的问题
	   时间复杂度 O(n) 空间复杂度为O（1）
	*/
	sum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
			customers[i] = 0
		}
	}

	num, maxVal := customers[0], customers[0]
	for i := 1; i < len(customers); i++ {
		if i < X {
			num += customers[i]
		} else {
			num += customers[i] - customers[i-X]
		}
		maxVal = max(maxVal, num)
	}
	return sum + maxVal
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
