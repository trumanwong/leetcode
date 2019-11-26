package minMoves

import "math"

func MinMoves(nums []int) int {
	min := math.MaxInt32
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	res := 0
	for _, v := range nums {
		res += v - min
	}
	return res
}