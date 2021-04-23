package numMovesStones

import "sort"

func NumMovesStones(a int, b int, c int) []int {
	nums := []int{a, b, c}
	sort.Ints(nums)
	a, b, c = nums[0], nums[1], nums[2]
	max := c - b - 1 + b - a - 1
	min := 0
	if c-b != 1 {
		min = 1
	}
	if b-a != 1 {
		min += 1
	}
	if b-a == 2 || c-b == 2 {
		min = 1
	}
	return []int{min, max}
}
