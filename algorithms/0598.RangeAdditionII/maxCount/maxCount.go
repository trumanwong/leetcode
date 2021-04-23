package maxCount

import "math"

func MaxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		m = int(math.Min(float64(m), float64(op[0])))
		n = int(math.Min(float64(n), float64(op[1])))
	}
	return m * n
}
