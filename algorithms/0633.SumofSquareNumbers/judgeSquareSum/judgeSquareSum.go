package judgeSquareSum

import "math"

func JudgeSquareSum(c int) bool {
	start, end := 0, int(math.Sqrt(float64(c)))
	for start <= end {
		result := start * start + end * end
		if result == c {
			return true
		} else if result > c {
			end--
		} else {
			start++
		}
	}
	return false
}