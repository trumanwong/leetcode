package totalMoney

func totalMoney(n int) int {
	res, step, start := 0, 1, 0
	for i := 1; i <= n; i++ {
		res += step
		if i%7 == 0 {
			step = start + 1
			start++
		}
		step++
	}
	return res
}
