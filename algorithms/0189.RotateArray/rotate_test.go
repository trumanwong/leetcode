package main

import (
	"leetcode/algorithms/0189.RotateArray/rotate"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, test := range tests {
		temp := []int{}
		temp = append(temp, test.input...)
		rotate.Rotate(temp, test.k)
		for i, v := range temp {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v, k = %d; expected %v", temp, test.input, test.k, test.output)
			}
		}

	}
}
