package rob

func Rob(nums []int) int {
	last, res := 0, 0
	for _, v := range nums {
		temp := last
		last = res
		if res < temp+v {
			res = temp + v
		}
	}
	return res
}
