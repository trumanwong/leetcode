package maxNumberOfApples

import "sort"

func maxNumberOfApples(arr []int) int {
	sort.Ints(arr)
	ret, sum := 0, 0
	for _, v := range arr {
		sum += v
		if sum <= 5000 {
			ret++
		} else {
			break
		}
	}
	return ret
}