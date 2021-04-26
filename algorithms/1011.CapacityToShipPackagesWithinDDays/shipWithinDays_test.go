package main

import (
	"leetcode/algorithms/1011.CapacityToShipPackagesWithinDDays/shipWithinDays"
	"testing"
)

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		weights []int
		D       int
		output  int
	}{
		{[]int{1,2,3,4,5,6,7,8,9,10}, 5, 15},
		{[]int{3,2,2,4,1,4}, 3, 6},
		{[]int{1,2,3,1,1}, 4, 3},
	}

	for _, test := range tests {
		ret := shipWithinDays.ShipWithinDays(test.weights, test.D)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.weights, test.D, test.output)
		}
	}
}
