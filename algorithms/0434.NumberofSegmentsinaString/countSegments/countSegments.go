package countSegments

func CountSegments(s string) int {
	res, letters := 0, 0
	for _, v := range s {
		if v != ' ' {
			letters++
		} else {
			if letters > 0 {
				res++
			}
			letters = 0
		}
	}
	if letters > 0 {
		res++
	}
	return res
}
