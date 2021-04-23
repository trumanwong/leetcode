package totalHammingDistance

func TotalHammingDistance(nums []int) int {
	res, bitcnt := 0, make([]int, 32)
	for _, v := range nums {
		for j := 0; j < 32; j++ {
			bitcnt[j] += v & 1
			v = v >> 1
		}
	}

	for i := 0; i < 32; i++ {
		res += bitcnt[i] * (len(nums) - bitcnt[i])
	}
	return res
}
