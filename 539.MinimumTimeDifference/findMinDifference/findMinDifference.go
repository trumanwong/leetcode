package findMinDifference

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	times := make([]int, 0)

	for _, timePoint := range timePoints {
		hour, _ := strconv.Atoi(string(timePoint[0:2]))
		minute, _ := strconv.Atoi(string(timePoint[3:]))

		timestamp := hour * 60 + minute
		times = append(times, timestamp)
	}

	sort.Ints(times)

	times = append(times, times[0] + 1440)
	ret := 1440

	for i := 1; i < len(times); i++ {
		ret = min(ret, times[i] - times[i - 1])
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}