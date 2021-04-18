package main

import (
	"leetcode/algorithms/5735.MaximumIceCreamBars/maxIceCream"
	"testing"
)

func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		costs  []int
		coins  int
		output int
	}{
		{[]int{1, 3, 2, 4, 1}, 7, 4},
		{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
		{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
	}

	for _, test := range tests {
		ret := maxIceCream.MaxIceCream(test.costs, test.coins)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.costs, test.coins, test.output)
		}
	}
}
