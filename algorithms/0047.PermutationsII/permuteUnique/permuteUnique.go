package permuteUnique

import "sort"

func PermuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	var perm func(arr []int, start int)
	perm = func(arr []int, start int) {
		if start == len(arr) - 1 {
			t := make([]int, len(arr))
			copy(t, arr)
			res = append(res, t)
			return
		}

		temp := make([]int, len(arr))
		copy(temp, arr)
		for i := start; i < len(arr); i++ {
			if i != start && temp[start] == temp[i] {
				continue // 去重
			}

			temp[start], temp[i] = temp[i], temp[start]
			perm(temp, start + 1)
		}
	}

	perm(nums, 0)
	return res
}