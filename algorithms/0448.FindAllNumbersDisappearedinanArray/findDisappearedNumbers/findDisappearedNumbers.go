package findDisappearedNumbers

func FindDisappearedNumbers(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
