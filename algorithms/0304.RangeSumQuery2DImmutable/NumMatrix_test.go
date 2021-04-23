package main

import (
	numMatrix "leetcode/algorithms/0304.RangeSumQuery2DImmutable/NumMatrix"
	"testing"
)

func TestNumMatrix(t *testing.T) {
	matrix := numMatrix.Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	ret := matrix.SumRegion(2, 1, 4, 3)
	if ret != 8 {
		t.Errorf("Got %d; expected 8", ret)
	}
	ret = matrix.SumRegion(1, 1, 2, 2)
	if ret != 11 {
		t.Errorf("Got %d; expected 11", ret)
	}
	ret = matrix.SumRegion(1, 2, 2, 4)
	if ret != 12 {
		t.Errorf("Got %d; expected 12", ret)
	}
}
