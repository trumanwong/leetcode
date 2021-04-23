package subarraySum

func SubarraySum(nums []int, k int) int {
	ret, index := 0, 0

	for index < len(nums) {
		sum := 0
		for i := index; i < len(nums); i++ {
			sum += nums[i]

			if sum == k {
				ret++
			}
		}
		index++
	}
	return ret
}
