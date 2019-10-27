package main

import (
	"leetcode/algorithms/0806.NumberofLinesToWriteStringWords/numberOfLines"
	"testing"
)

func TestNumberOfLines(t *testing.T)  {
	tests := []struct{
		widths []int
		S string
		output []int
	}{
		{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"abcdefghijklmnopqrstuvwxyz", []int{3, 60}},
		{[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"bbbcccdddaaa", []int{2, 4}},
	}

	for _, test := range tests {
		ret := numberOfLines.NumberOfLines(test.widths, test.S)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for widths = %v, S = %s; expected %v", ret, test.widths, test.S, test.output)
			}
		}
	}
}