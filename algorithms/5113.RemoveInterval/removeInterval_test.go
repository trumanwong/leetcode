package main

import (
	"leetcode/algorithms/5113.RemoveInterval/removeInterval"
	"reflect"
	"testing"
)

func TestRemoveInterval(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		toBeRemoved []int
		output      [][]int
	}{
		{[][]int{{0, 2}, {3, 4}, {5, 7}}, []int{1, 6}, [][]int{{0, 1}, {6, 7}}},
		{[][]int{{0, 5}}, []int{2, 3}, [][]int{{0, 2}, {3, 5}}},
	}

	for _, test := range tests {
		ret := removeInterval.RemoveInterval(test.intervals, test.toBeRemoved)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for intervals = %v, toBeRemoved = %v; expected %v", ret, test.intervals, test.toBeRemoved, test.output)
		}
	}
}
