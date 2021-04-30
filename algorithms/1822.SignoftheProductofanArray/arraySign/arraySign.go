package arraySign

func arraySign(nums []int) int {
	result := 1
	for _, v := range nums {
		if v == 0 {
			result = 0
			break
		}
		if v < 0 {
			result *= -1
		}
	}
	return result
}