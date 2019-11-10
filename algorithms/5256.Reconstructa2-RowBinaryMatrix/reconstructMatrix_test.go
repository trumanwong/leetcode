package main

import (
	"leetcode/algorithms/5256.Reconstructa2-RowBinaryMatrix/reconstructMatrix"
	"reflect"
	"testing"
)

func TestReconstructMatrix(t *testing.T)  {
	tests := []struct{
		upper int
		lower int
		colsum []int
		output [][]int
	} {
		{2, 1, []int{1, 1, 1}, [][]int{{0, 1, 1}, {1, 0, 0}}},
		{2, 3, []int{2,2,1,1}, [][]int{}},
		{5,5,[]int{2,1,2,0,1,0,1,2,0,1}, [][]int{{1,1,1,0,1,0,0,1,0,0}, {1,0,1,0,0,0,1,1,0,1}}},
	}

	for _, test := range tests {
		ret := reconstructMatrix.ReconstructMatrix(test.upper, test.lower, test.colsum)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for upper = %d, lower = %d, colsum = %v; expected %v", ret, test.upper, test.lower, test.colsum, test.output)
		}
	}
}