package kConcatenationMaxSum

import "math"

func kConcatenationMaxSum(arr []int, k int) int {
	sum, arrSum, ret := 0, 0, 0
	mod := int(math.Pow10(9)) + 7
	for i := 0; i < len(arr); i++ {
		arrSum = arrSum + arr[i]
	}
	for i := 0; i < k; i++ {
		if i == 0 || i == 1 || i == k-2 || i == k-1 {
			for index := 0; index < len(arr); index++ {
				sum = sum + arr[index]
				if sum < 0 {
					sum = 0
				} else {
					ret = max(ret, sum)
				}
			}
		} else {
			sum = sum + arrSum
			if sum < 0 {
				sum = 0
			} else {
				ret = max(ret, sum)
			}
		}
	}
	ret = max(0, ret)
	return ret % mod
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
