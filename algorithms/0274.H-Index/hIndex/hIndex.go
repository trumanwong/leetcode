package hIndex

import "sort"

func HIndex(citations []int) int {
	sort.Ints(citations)
	ret := len(citations)
	for _, v := range citations {
		if v >= ret {
			return ret
		} else {
			ret--
		}
	}
	return ret
}
