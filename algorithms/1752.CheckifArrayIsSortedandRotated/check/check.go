package check

func check(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 0 || nums[0] >= nums[len(nums)-1] {
		return true
	}
	return false
}
