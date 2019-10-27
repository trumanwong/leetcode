package missingNumber

func MissingNumber(nums []int) int {
	sum := len(nums)
	for i := 0; i < len(nums); i++ {
		sum ^= nums[i]
		sum ^= i
	}
	return sum
}