package findMin

import "math"

func FindMin(nums []int) int {
	res := math.MaxInt32
	for _, v := range nums {
		if v < res {
			res = v
		}
	}
	return res
}