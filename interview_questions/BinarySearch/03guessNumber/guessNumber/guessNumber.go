package guessNumber

import "math/rand"

func guess(num int) int {
	pick := rand.Int()
	if pick > num {
		return 1
	} else if pick < num {
		return -1
	}
	return 0
}

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		result := guess(mid)
		if result == 0 {
			return mid
		} else if result == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0
}