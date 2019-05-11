package countPrimeSetBits

func CountPrimeSetBits(L int, R int) int {
	count := 0
	for L <= R {
		temp, binnary := L, 0
		for temp != 0 {
			if temp & 1 == 1 {
				binnary++
			}
			temp >>= 1
		}
		if isPrime(binnary) {
			count++
		}
		L++
	}
	return count
}

func isPrime(num int) bool {
	return num == 2||num == 3||num == 5||num == 7||num == 11||num == 13||num == 17||num == 19
}