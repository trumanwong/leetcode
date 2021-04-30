package missingNumber

func missingNumber(nums []int) int {
	res := 0
	for i := 1; i <= len(nums); i++ {
		res = res ^ i
		res = res ^ nums[i - 1]
	}
	return res
}