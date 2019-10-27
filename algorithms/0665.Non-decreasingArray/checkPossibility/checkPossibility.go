package checkPossibility

func CheckPossibility(nums []int) bool {
	count, end := 0, len(nums) - 1
	if end < 2 {
		return true
	}
	if nums[0] > nums[1] {
		nums[0] = nums[1]
		count++
	}
	for i := 1; i < end; i++ {
		right := nums[i + 1]
		if nums[i] > right {
			count++
			if count > 1 {
				return false
			}
			left := nums[i - 1]
			if left > right {
				nums[i + 1] = nums[i]
			} else {
				nums[i] = left
			}
		}
	}
	return true
}