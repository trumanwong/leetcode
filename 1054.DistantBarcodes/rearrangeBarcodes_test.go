package main

import (
	"fmt"
	"testing"
	"truman.com/leetcode/1054.DistantBarcodes/rearrangeBarcodes"
)

func TestRearrangeBarcodes(t *testing.T)  {
	tests := []struct{
		barcodes []int
		output []int
	}{
		{[]int{1,1,1,2,2,2}, []int{2,1,2,1,2,1}},
		{[]int{1,1,1,1,2,2,3,3}, []int{1,3,1,3,2,1,2,1}},
		{[]int{2,1,1}, []int{1,2,1}},
		{[]int{2,1,1,2}, []int{2,1,2,1}},
		{[]int{1,2,2,2,1},[]int{2,1,2,1,2}},
	}

	for _, test := range tests {
		ret := rearrangeBarcodes.RearrangeBarcodes(test.barcodes)
		fmt.Println(ret)
	}
}