package merge

import "sort"

func Merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	starts, ends := make([]int, 0), make([]int, 0)
	for i := 0; i < len(intervals); i++ {
		starts = append(starts, intervals[i][0])
		ends = append(ends, intervals[i][1])
	}

	sort.Ints(starts)
	sort.Ints(ends)
	for i, j := 0, 0; i < len(intervals); i++ {
		if i == len(intervals) - 1 || starts[i + 1] > ends[i] {
			res = append(res, []int{starts[j], ends[i]})
			j = i + 1
		}
	}
	return res
}