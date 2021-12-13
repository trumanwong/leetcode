package maximumWealth

import "math"

func maximumWealth(accounts [][]int) int {
	res, sum := 0, 0
	for _, account := range accounts {
		sum = 0
		for _, v := range account {
			sum += v
		}
		res = int(math.Max(float64(res), float64(sum)))
	}
	return res
}