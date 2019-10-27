package heightChecker

import "sort"

func HeightChecker(heights []int) int {
	copyData := make([]int, len(heights))
	copy(copyData, heights)
	sort.Ints(copyData)

	res := 0
	for i, v := range copyData {
		if v != heights[i] {
			res++
		}
	}
	return res
}