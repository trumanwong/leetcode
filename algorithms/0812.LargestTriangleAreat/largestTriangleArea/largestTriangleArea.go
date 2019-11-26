package largestTriangleArea

import (
	"math"
)

func LargestTriangleArea(points [][]int) float64 {
	var res float64
	for i := 0; i < len(points) ; i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				res = math.Max(res, area(points[i], points[j], points[k]))
			}
		}
	}
	return res
}

func area(P, Q, R []int) float64 {
	return 0.5 * math.Abs(float64(P[0]*Q[1] + Q[0]*R[1] + R[0]*P[1] - P[1]*Q[0] - Q[1]*R[0] - R[1]*P[0]))
}