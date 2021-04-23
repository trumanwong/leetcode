package maximumProduct

import "math"

func MaximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v > max1 {
			max1, max2, max3 = v, max1, max2
		} else if v > max2 {
			max2, max3 = v, max2
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min1, min2 = v, min1
		} else if v < min2 {
			min2 = v
		}
	}
	max := max1 * max2 * max3
	min := max1 * min1 * min2
	if max > min {
		return max
	}
	return min
}
