package main

import (
	"leetcode/algorithms/1054.DistantBarcodes/rearrangeBarcodes"
	"testing"
)

func TestRearrangeBarcodes(t *testing.T)  {
	tests := []struct{
		barcodes []int
		output []int
	}{
		{[]int{1,1,1,2,2,2}, []int{1, 2, 1, 2, 1, 2}},
		{[]int{1,1,1,1,2,2,3,3}, []int{1, 2, 1, 2, 1, 3, 1, 3}},
		{[]int{2,1,1}, []int{1,2,1}},
		{[]int{2,1,1,2}, []int{2,1,2,1}},
		{[]int{1,2,2,2,1},[]int{2,1,2,1,2}},
	}

	for _, test := range tests {
		ret := rearrangeBarcodes.RearrangeBarcodes(test.barcodes)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.barcodes, test.output)
				break
			}
		}
	}
}