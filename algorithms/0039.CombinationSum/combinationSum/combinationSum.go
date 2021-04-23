package combinationSum

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	var backtrack func(i int, tmp_sum int, tmp []int)
	backtrack = func(i int, tmp_sum int, tmp []int) {
		if tmp_sum == target {
			temp := make([]int, len(tmp))
			copy(temp, tmp)
			res = append(res, temp)
			return
		}
		for j := i; j < len(candidates); j++ {
			if tmp_sum+candidates[j] > target {
				break
			}
			backtrack(j, tmp_sum+candidates[j], append(tmp, candidates[j]))
		}
	}

	backtrack(0, 0, []int{})
	return res
}
