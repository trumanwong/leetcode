package checkStraightLine

import "math"

func CheckStraightLine(coordinates [][]int) bool {
	k := float64(coordinates[1][1]-coordinates[0][1]) / float64(coordinates[1][0]-coordinates[0][0])
	b := float64(coordinates[0][1]) - k*float64(coordinates[0][0])
	for _, v := range coordinates {
		if math.Abs(float64(v[1])-float64(v[0])*k-b) > 0.0001 {
			return false
		}
	}
	return true
}
