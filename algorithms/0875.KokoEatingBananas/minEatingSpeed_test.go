package main

import (
	"leetcode/algorithms/0875.KokoEatingBananas/minEatingSpeed"
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		piles  []int
		H      int
		output int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}

	for _, test := range tests {
		ret := minEatingSpeed.MinEatingSpeed(test.piles, test.H)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.piles, test.H, test.output)
		}
	}
}
