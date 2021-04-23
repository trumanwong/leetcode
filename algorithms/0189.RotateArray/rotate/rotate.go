package rotate

func Rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	start, k := 0, k%n
	for k != 0 {
		for i := 0; i < k; i++ {
			nums[start], nums[n-k+start] = nums[n-k+start], nums[start]
			start++
		}
		n = n - k
		k %= n
	}
}
