package main

import (
	"leetcode/algorithms/1260.Shift2DGrid/shiftGrid"
	"reflect"
	"testing"
)

func TestShiftGrid(t *testing.T)  {
	tests := []struct{
		grid [][]int
		grid2 [][]int
		k int
		output [][]int
	}{
		{[][]int{{1,2,3},{4,5,6},{7,8,9}},[][]int{{1,2,3},{4,5,6},{7,8,9}}, 1,
		[][]int{{9,1,2},{3,4,5},{6,7,8}}},

		{[][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}},
			[][]int{{3,8,1,9},{19,7,2,5},{4,6,11,10},{12,0,21,13}},4,
		[][]int{{12,0,21,13},{3,8,1,9},{19,7,2,5},{4,6,11,10}}},

		{[][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{1,2,3},{4,5,6},{7,8,9}}, 9,
		[][]int{{1,2,3},{4,5,6},{7,8,9}}},
	}

	for _, test := range tests {
		ret := shiftGrid.ShiftGrid(test.grid, test.k)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.grid2, test.output)
		}
	}
}