package maxProfit

func MaxProfit(prices []int) int {
	max := 0
	length := len(prices)
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}
