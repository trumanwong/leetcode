package hIndex

func HIndex(citations []int) int {
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