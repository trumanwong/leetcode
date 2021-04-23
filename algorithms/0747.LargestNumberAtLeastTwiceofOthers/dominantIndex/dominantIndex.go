package dominantIndex

func DominantIndex(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	first, second, index := -1, -1, -1
	for i, v := range nums {
		if v > first {
			index = i
			first, second = v, first
		} else if v > second {
			second = v
		}
	}
	if first < second*2 {
		return -1
	}
	return index
}
