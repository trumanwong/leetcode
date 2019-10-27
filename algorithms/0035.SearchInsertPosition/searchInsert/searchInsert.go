package searchInsert

func SearchInsert(nums []int, target int) int {
	for i, v := range nums {
		if target <= v {
			return i
		}
	}
	return len(nums)
}
