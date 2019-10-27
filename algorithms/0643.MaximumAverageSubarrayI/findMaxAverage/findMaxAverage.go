package findMaxAverage

func FindMaxAverage(nums []int, k int) float64 {
	var max float64
	sum, max := 0, -10000000
	count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count++
		if count >= k {
			if count > k {
				sum -= nums[i - k]
				count--
			}
			if float64(sum) > max {
				max = float64(sum)
			}
		}
	}
	return max / float64(k)
}