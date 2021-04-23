package minMoves2

import (
	"math"
	"sort"
)

func MinMoves2(nums []int) int {
	sort.Ints(nums)
	res := 0
	for _, v := range nums {
		res += int(math.Abs(float64(nums[len(nums)/2] - v)))
	}
	return res
}
