package permute

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	visited := make([]int, len(nums))
	var backtrack func(tmp []int)
	backtrack = func(tmp []int) {
		if len(tmp) == len(nums) {
			temp := make([]int, len(tmp))
			copy(temp, tmp)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] == 1 {
				continue
			}

			visited[i] = 1;
			tmp = append(tmp, nums[i])
			backtrack(tmp)
			visited[i] = 0
			tmp = tmp[:len(tmp) - 1]
		}
	}
	backtrack([]int{})
	return res
}