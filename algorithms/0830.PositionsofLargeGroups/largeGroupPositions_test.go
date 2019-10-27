package main

import (
	"leetcode/algorithms/0830.PositionsofLargeGroups/largeGroupPositions"
	"testing"
)

func TestLargeGroupPositions(t *testing.T)  {
	tests := []struct{
		S string
		output [][]int
	}{
		{"abbxxxxzzy", [][]int{{3,6}}},
		{"abc", [][]int{}},
		{"abcdddeeeeaabbbcd", [][]int{{3,5}, {6, 9}, {12,14}}},
	}

	for _, test := range tests {
		ret := largeGroupPositions.LargeGroupPositions(test.S)
		for i, arr := range ret {
			for j, v := range arr {
				if test.output[i][j] != v {
					t.Errorf("Got %v for input %s; expected %v", ret, test.S, test.output)
					break
				}
			}
		}
	}
}