package checkOnesSegment

func checkOnesSegment(s string) bool {
	res, judge := true, true
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			if !judge {
				res = false
			}
		} else {
			judge = false
		}
	}
	return res
}