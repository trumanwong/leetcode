package main

import (
	"leetcode/algorithms/0849.MaximizeDistancetoClosestPerson/maxDistToClosest"
	"testing"
)

func TestMaxDistToClosest(t *testing.T)  {
	tests := []struct{
		seats []int
		output int
	}{
		{[]int{1, 0, 0, 0, 1, 0, 1}, 2},
		{[]int{1, 0, 0, 0}, 3},
	}

	for _, test := range tests {
		ret := maxDistToClosest.MaxDistToClosest(test.seats)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.seats, test.output)
		}
	}
}