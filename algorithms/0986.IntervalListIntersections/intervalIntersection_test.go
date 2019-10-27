package main

import (
	"leetcode/algorithms/0986.IntervalListIntersections/intervalIntersection"
	"testing"
)

func TestIntervalIntersection(t *testing.T)  {
	tests := []struct{
		A [][]int
		B [][]int
		output [][]int
	}{
		{[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}},
	}

	for _, test := range tests {
		ret := intervalIntersection.IntervalIntersection(test.A, test.B)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for input A = %v, B = %v; expected %v", ret, test.A, test.B, test.output)
				break
			}
		}
	}
}