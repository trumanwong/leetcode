package largestAltitude

func largestAltitude(gain []int) int {
	res, level := 0, 0
	for _, v := range gain {
		level += v
		if level > res {
			res = level
		}
	}
	return res
}