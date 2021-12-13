package getMinDistance

func getMinDistance(nums []int, target int, start int) int {
	for i := range nums {
		if (start + i < len(nums) && nums[i + start] == target) || (start - i >= 0 && nums[start - i] == target) {
			return i
		}
		if start + i >= len(nums) && start - i < 0 {
			break
		}
	}
	return -1
}