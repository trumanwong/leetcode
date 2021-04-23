package subsetsWithDup

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	var dfs func([]int, int)
	dfs = func(temp []int, index int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := index; i < len(nums); i++ {
			if i != index && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
