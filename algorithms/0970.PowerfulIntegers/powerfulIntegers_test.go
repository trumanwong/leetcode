package main

import (
	"leetcode/algorithms/0970.PowerfulIntegers/powerfulIntegers"
	"testing"
)

func TestPowerfulIntegers(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		bound  int
		output []int
	}{
		{2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10}},
		{3, 5, 15, []int{2, 4, 6, 8, 10, 14}},
	}

	for _, test := range tests {
		ret := powerfulIntegers.PowerfulIntegers(test.x, test.y, test.bound)
		for i, v := range test.output {
			if v != test.output[i] {
				t.Errorf("Got %v for input x = %d, y = %d, bound = %d; expected %v", ret, test.x, test.y, test.bound, test.output)
				break
			}
		}
	}
}
