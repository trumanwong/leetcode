package findKthLargest

func FindKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums[k-1]
}
