package combinationSum2

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	var backtrack func(i int, sum int, tmp []int)
	backtrack = func(i int, sum int, tmp []int) {
		if sum == target {
			temp := make([]int, len(tmp))
			copy(temp, tmp)
			res = append(res, temp)
			return
		}

		for j := i; j < len(candidates); j++ {
			if sum+candidates[j] > target {
				break
			}
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			backtrack(j+1, sum+candidates[j], append(tmp, candidates[j]))
		}
	}
	backtrack(0, 0, []int{})
	return res
}
