package constructRectangle

import "math"

func ConstructRectangle(area int) []int {
	width := int(math.Sqrt(float64(area)))
	for area%width != 0 {
		width--
	}
	return []int{area / width, width}
}
