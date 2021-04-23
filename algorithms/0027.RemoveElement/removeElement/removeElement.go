package removeElement

func RemoveElement(nums []int, val int) int {
	i := 0
	length := len(nums)
	for i < length {
		if nums[i] == val {
			nums[i] = nums[length-1]
			length = length - 1
		} else {
			i++
		}
	}
	return length
}
