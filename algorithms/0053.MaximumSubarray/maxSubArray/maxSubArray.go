package maxSubArray

func MaxSubArray(nums []int) int {
	res, sum := nums[0], 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > res {
			res = sum
		}
	}
	return res
}
