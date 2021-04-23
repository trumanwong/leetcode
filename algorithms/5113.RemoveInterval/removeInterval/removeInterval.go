package removeInterval

func RemoveInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := make([][]int, 0)

	for _, interval := range intervals {
		if interval[0] > toBeRemoved[0] && interval[1] < toBeRemoved[1] {
			continue
		} else if interval[0] < toBeRemoved[0] {
			if interval[1] < toBeRemoved[0] {
				res = append(res, interval)
			} else {
				res = append(res, []int{interval[0], toBeRemoved[0]})
			}

			if interval[1] > toBeRemoved[1] {
				res = append(res, []int{toBeRemoved[1], interval[1]})
			}
		} else if interval[1] > toBeRemoved[1] {
			if interval[0] > toBeRemoved[1] {
				res = append(res, interval)
			} else {
				res = append(res, []int{toBeRemoved[1], interval[1]})
			}
		}
	}
	return res
}
