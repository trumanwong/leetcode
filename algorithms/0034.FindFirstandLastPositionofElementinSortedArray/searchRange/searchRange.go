package searchRange

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	left := 0
	right := len(nums) - 1
	for ; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			for ; mid >= left && nums[mid] == target; {
				mid--
			}
			result[0] = mid + 1
			mid = (left + right) / 2
			for ; mid <= right && nums[mid] == target; {
				mid++
			}
			result[1] = mid - 1
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}