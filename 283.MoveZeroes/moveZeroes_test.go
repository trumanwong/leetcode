package main

import (
	"testing"
	"truman.com/leetcode/283.MoveZeroes/moveZeroes"
)

func TestMoveZeroes(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{0,1,0,3,12},[]int{1,3,12,0,0}},
		{[]int{0},[]int{0}},
	}
	for _, test := range tests {
		temp := make([]int, len(test.input))
		copy(temp, test.input)
		moveZeroes.MoveZeroes(temp)
		for i, v := range temp {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", temp, test.input, test.output)
				break
			}
		}
	}
}