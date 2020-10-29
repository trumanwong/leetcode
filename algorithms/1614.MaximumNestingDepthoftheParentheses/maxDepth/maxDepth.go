package maxDepth

func MaxDepth(s string) int {
	res, curr := 0, 0
	parentheses := []byte("()")
	for _, v := range s {
		if byte(v) == parentheses[0] {
			curr++
		} else if byte(v) == parentheses[1] {
			if curr > res {
				res = curr
			}
			curr--
		}
	}
	return res
}
