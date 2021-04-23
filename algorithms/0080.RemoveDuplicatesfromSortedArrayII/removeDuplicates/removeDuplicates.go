package removeDuplicates

func RemoveDuplicates(nums []int) int {
	res := 0
	for _, v := range nums {
		if res < 2 || v != nums[res-2] {
			nums[res] = v
			res++
		}
	}
	return res
}
