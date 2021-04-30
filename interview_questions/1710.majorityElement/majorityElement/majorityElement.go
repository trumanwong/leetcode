package majorityElement

func majorityElement(nums []int) int {
	x, votes := 0, 0
	for _, v := range nums {
		if votes == 0 {
			x = v
		}

		if x == v {
			votes++
		} else {
			votes--
		}
	}
	count := 0
	for _, v := range nums {
		if v == x {
			count++
		}
	}
	if count > len(nums)/2 {
		return x
	}
	return -1
}
