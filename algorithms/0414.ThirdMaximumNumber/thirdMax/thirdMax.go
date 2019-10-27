package thirdMax

import (
	"math"
)

func ThirdMax(nums []int) int {
	first, second, third := math.MinInt32, math.MinInt32, math.MinInt32
	hasMin := false
	for _, num := range nums {
		if num == math.MinInt32 {
			hasMin = true
		}
		if num > first {
			first, second, third = num, first, second
		} else if num == first {
			continue
		} else if num > second {
			second, third = num, second
		} else if num == second {
			continue
		} else if num > third {
			third = num
		}
	}
	if (!hasMin && third == math.MinInt32) || (hasMin && third == second ){
		return first
	}
	return third
}