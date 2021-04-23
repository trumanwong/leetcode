package containsNearbyAlmostDuplicate

func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
