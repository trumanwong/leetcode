package main

import (
	"testing"
	"truman.com/leetcode/rotate/rotate"
)

func TestRotate(t *testing.T)  {
	input1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	output1 := [][]int{
		{7,4,1},
		{8,5,2},
		{9,6,3},
	}
	tmp := input1
	rotate.Rotate(input1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if input1[i][j] != output1[i][j] {
				t.Errorf("Got %v for %v;expected %v", input1, tmp, output1)
			}
		}
	}
	input2 := [][]int{
		{5, 1, 9,11},
		{2, 4, 8,10},
		{13, 3, 6, 7},
		{15,14,12,16},
	}
	output2 := [][]int{
		{15,13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7,10,11},
	}
	tmp = input2
	rotate.Rotate(input2)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if input2[i][j] != output2[i][j] {
				t.Errorf("Got %v for %v;expected %v", input2, tmp, output2)
			}
		}
	}
}
