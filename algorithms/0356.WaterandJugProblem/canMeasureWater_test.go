package main

import (
	"leetcode/algorithms/0356.WaterandJugProblem/canMeasureWater"
	"testing"
)

func TestCanMeasureWater(t *testing.T)  {
	tests := []struct{
		x int
		y int
		z int
		output bool
	}{
		{3, 5, 4, true},
		{2, 6, 5, false},
	}

	for _, test := range tests {
		ret := canMeasureWater.CanMeasureWater(test.x, test.y, test.z)
		if ret != test.output {
			t.Errorf("Got %t for input %d, %d, %d; expected %t", ret, test.x, test.y, test.z, test.output)
		}
	}
}