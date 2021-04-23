package uncommonFromSentences

func UncommonFromSentences(A string, B string) []string {
	m := make(map[string]int)
	isFirst, start, last := true, 0, 0
	for i, v := range A {
		if isFirst && v != ' ' {
			start, isFirst = i, false
		}
		if !isFirst && v == ' ' {
			m[string(A[start:last])]++
			isFirst = true
		}
		last++
	}
	if !isFirst {
		m[string(A[start:last])]++
	}

	start, last = 0, 0
	for i, v := range B {
		if isFirst && v != ' ' {
			start, isFirst = i, false
		}
		if !isFirst && v == ' ' {
			m[string(B[start:last])]++
			isFirst = true
		}
		last++
	}
	if !isFirst {
		m[string(B[start:last])]++
	}
	res := []string{}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
