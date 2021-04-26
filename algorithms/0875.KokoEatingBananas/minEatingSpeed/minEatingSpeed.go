package minEatingSpeed

import "math"

func MinEatingSpeed(piles []int, h int) int {
	lo, hi := 1, int(math.Pow(10, 9))
	for lo < hi {
		mi := lo + (hi-lo)/2
		if !possible(piles, h, mi) {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return hi
}

func possible(piles []int, h, k int) bool {
	time := 0
	for _, v := range piles {
		time += (v-1)/k + 1
	}
	return time <= h
}
