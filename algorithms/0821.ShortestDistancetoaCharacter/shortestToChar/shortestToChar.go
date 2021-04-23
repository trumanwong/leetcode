package shortestToChar

func ShortestToChar(S string, C byte) []int {
	eArr := []int{}
	for i, v := range S {
		if byte(v) == C {
			eArr = append(eArr, i)
		}
	}

	res := make([]int, len(S))
	for i, v := range S {
		if byte(v) == C {
			continue
		}
		res[i] = minDiff(i, eArr, len(S))
	}
	return res
}

func minDiff(index int, arr []int, minVal int) int {
	for _, v := range arr {
		diff := abs(index - v)
		if diff < minVal {
			minVal = diff
		}
	}
	return minVal
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
