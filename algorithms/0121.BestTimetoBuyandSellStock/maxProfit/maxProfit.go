package maxProfit

import "math"

func MaxProfit(prices []int) int {
	max := 0
	min := math.MaxInt32
	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > max {
			max = price - min
		}
	}
	return max
}
