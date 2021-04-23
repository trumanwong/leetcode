package movesToMakeZigzag

func MovesToMakeZigzag(nums []int) int {
	odd, even := 0, 0
	nums_copy := make([]int, len(nums))
	copy(nums_copy, nums)
	i := 1
	for i < len(nums) {
		for nums[i] >= nums[i-1] {
			nums[i]--
			even++
		}
		for i+1 < len(nums) && nums[i] >= nums[i+1] {
			nums[i]--
			even++
		}
		i += 2
	}
	i = 0
	for i < len(nums) {
		for i+1 < len(nums) && nums_copy[i] >= nums_copy[i+1] {
			nums_copy[i]--
			odd++
		}
		for i-1 >= 0 && nums_copy[i] >= nums_copy[i-1] {
			nums_copy[i]--
			odd++
		}
		i += 2
	}
	if even > odd {
		return odd
	}
	return even
}
