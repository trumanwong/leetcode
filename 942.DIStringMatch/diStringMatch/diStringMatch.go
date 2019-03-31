package diStringMatch

func DiStringMatch(S string) []int {
	N := len(S)
	start := 0
	res := make([]int, N + 1)
	last := 0
	for i, v := range S {
		if string(v) == "I" {
			res[i] = start
			start++
			last = start
		} else {
			res[i] = N
			N--
			last = N
		}
	}
	res[len(S)] = last
	return res
}