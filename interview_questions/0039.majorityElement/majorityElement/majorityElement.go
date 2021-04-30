package majorityElement

func majorityElement(nums []int) int {
	x, votes := 0, 0
	for _, v := range nums {
		if votes == 0 {
			x = v
		}
		if v == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}