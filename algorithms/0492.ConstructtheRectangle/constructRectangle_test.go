package main

import (
	"leetcode/algorithms/0492.ConstructtheRectangle/constructRectangle"
	"testing"
)

func TestConstructRectangle(t *testing.T) {
	tests := []struct {
		area   int
		output []int
	}{
		{4, []int{2, 2}},
	}

	for _, test := range tests {
		ret := constructRectangle.ConstructRectangle(test.area)
		for i, v := range ret {
			if test.output[i] != v {
				t.Errorf("Got %v for input %d; expected %v", ret, test.area, test.output)
				break
			}
		}
	}
}
