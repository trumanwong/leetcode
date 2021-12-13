package subsetXORSum

func subsetXORSum(nums []int) int {
	res, n := 0, len(nums)
	var dfs func(val, idx int)
	dfs = func(val, idx int) {
		if idx == n {
			res += val
			return
		}

		dfs(val^nums[idx], idx+1)
		dfs(val, idx+1)
	}
	dfs(0, 0)
	return res
}
