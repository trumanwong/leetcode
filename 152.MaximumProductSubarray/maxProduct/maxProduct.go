package maxProduct

func MaxProduct(nums []int) int {
	res, maxValue, minValue := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxValue, minValue = max(nums[i], nums[i] * minValue), min(nums[i], nums[i] * maxValue)
		} else {
			maxValue, minValue = max(nums[i], nums[i] * maxValue), min(nums[i], nums[i] * minValue)
		}
		res = max(res, maxValue)
	}
	return res
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}