package powerfulIntegers

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	if bound <= 1 {
		return []int{}
	}

	if x == 1 && y == 1 {
		return []int{2}
	}

	maps := make(map[int]int)
	for i := 0; math.Pow(float64(x), float64(i)) < float64(bound); i++ {
		for j := 0; math.Pow(float64(y), float64(j)) < float64(bound); j++ {
			temp := int(math.Pow(float64(x), float64(i)) + math.Pow(float64(y), float64(j)))
			if temp <= bound {
				maps[temp] = temp
			}
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	res := []int{}
	for _, v := range maps {
		res = append(res, v)
	}
	return res
}