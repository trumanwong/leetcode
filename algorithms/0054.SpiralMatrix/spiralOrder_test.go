package main

import (
	"leetcode/algorithms/0054.SpiralMatrix/spiralOrder"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		output []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, test := range tests {
		ret := spiralOrder.SpiralOrder(test.matrix)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.matrix, test.output)
				break
			}
		}
	}
}
