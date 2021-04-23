package maxDepthAfterSplit

func MaxDepthAfterSplit(seq string) []int {
	d := 0
	ret := make([]int, len(seq))
	for i, v := range seq {
		if v == '(' {
			ret[i] = d % 2
			d += 1
		} else {
			d -= 1
			ret[i] = d % 2
		}
	}
	return ret
}
