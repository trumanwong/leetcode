package findMaximumXOR

func FindMaximumXOR(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res = max(res, nums[i] ^ nums[j])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}