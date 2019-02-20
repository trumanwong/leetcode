package removeDuplicates

func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	i := 0
	for j := 1; j < length; j++{
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}