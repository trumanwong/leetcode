package secondHighest

import "strconv"

func secondHighest(s string) int {
	var max, second uint8
	max, second = 0, 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v >= '0' && v <= '9' {
			if v > max {
				second = max
				max = v
			} else if v < max && v > second {
				second = v
			}
		}
	}
	if second == 0 {
		return -1
	}
	res, _ := strconv.Atoi(string([]byte{second}))
	return res
}