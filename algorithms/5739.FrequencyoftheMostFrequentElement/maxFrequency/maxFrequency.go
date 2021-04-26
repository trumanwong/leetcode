package maxFrequency

import (
	"math"
	"sort"
)

func MaxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	var total int64
	l, res := 0, 1
	for r := 1; r < len(nums); r++ {
		total += int64((nums[r] - nums[r-1]) * (r - l))
		for total > int64(k) {
			total -= int64(nums[r] - nums[l])
			l++
		}
		res = int(math.Max(float64(res), float64(r-l+1)))
	}
	return res
}
