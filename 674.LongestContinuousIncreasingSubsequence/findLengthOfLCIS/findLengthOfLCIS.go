package findLengthOfLCIS

func FindLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	continuous, max, diff := 1, 0, nums[len(nums) - 1] - nums[len(nums) - 2]
	if diff <= 0 {
		max, continuous = 1, 1
	} else {
		continuous++
		max = continuous
	}
	for i := len(nums) - 2; i > 0; i-- {
		temp := nums[i] - nums[i - 1]
		if continuous > max {
			max = continuous
		}
		if temp <= 0 {
			diff, continuous = temp, 1
		} else {
			continuous++
		}
	}
	if continuous > max {
		max = continuous
	}
	return max
}