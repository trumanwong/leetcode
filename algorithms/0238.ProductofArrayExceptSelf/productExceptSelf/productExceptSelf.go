package productExceptSelf

func ProductExceptSelf(nums []int) []int {
	res, length := make([]int, len(nums)), len(nums)
	res[0] = 1
	for i := 0; i < length-1; i++ {
		res[i+1] = res[i] * nums[i]
	}
	mul := 1
	for i := length - 1; i >= 0; i-- {
		res[i] *= mul
		mul *= nums[i]
	}
	return res
}
