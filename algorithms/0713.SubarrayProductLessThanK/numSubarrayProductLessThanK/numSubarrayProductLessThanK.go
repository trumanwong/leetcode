package numSubarrayProductLessThanK

func NumSubarrayProductLessThanK(nums []int, k int) int {
	i, ret, mul, st := 0, 0, 1, 0
	for i < len(nums) {
		mul *= nums[i]
		for mul >= k && st <= i {
			mul /= nums[st]
			st++
		}
		ret += i - st + 1
		i++
	}
	return ret
}
