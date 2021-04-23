package findRelativeRanks

import (
	"sort"
	"strconv"
)

func FindRelativeRanks(nums []int) []string {
	temp, m, length := make([]int, len(nums)), make(map[int]int), len(nums)
	copy(temp, nums)
	sort.Ints(temp)
	for i, v := range temp {
		m[v] = length - i
	}
	res := make([]string, len(nums))
	for i, v := range nums {
		if m[v] == 1 {
			res[i] = "Gold Medal"
		} else if m[v] == 2 {
			res[i] = "Silver Medal"
		} else if m[v] == 3 {
			res[i] = "Bronze Medal"
		} else {
			res[i] = strconv.Itoa(m[v])
		}
	}
	return res
}
