package main

import (
	"leetcode/algorithms/1232.CheckIfItIsaStraightLine/checkStraightLine"
	"testing"
)

func TestCheckStraightLine(t *testing.T)  {
	tests := []struct{
		coordinates [][]int
		output bool
	}{
		{[][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}}, true},
		{[][]int{{1,1},{2,2},{3,4},{4,5},{5,6},{7,7}}, false},
	}

	for _, test := range tests {
		ret := checkStraightLine.CheckStraightLine(test.coordinates)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.coordinates, test.output)
		}
	}
}