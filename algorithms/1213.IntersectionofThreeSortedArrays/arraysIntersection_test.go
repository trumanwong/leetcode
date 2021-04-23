package main

import (
	"leetcode/algorithms/1213.IntersectionofThreeSortedArrays/arraysIntersection"
	"testing"
)

func TestArraysIntersection(t *testing.T) {
	tests := []struct {
		arr1   []int
		arr2   []int
		arr3   []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}, []int{1, 5}},
	}

	for _, test := range tests {
		res := arraysIntersection.ArraysIntersection(test.arr1, test.arr2, test.arr3)
		if len(res) != len(test.output) {
			t.Errorf("Got %v for input arr1:%v, arr2:%v, arr3:%v, expected %v", res, test.arr1, test.arr2, test.arr3, test.output)
		}

		for i, v := range res {
			if test.output[i] != v {
				t.Errorf("Got %v for input arr1:%v, arr2:%v, arr3:%v, expected %v", res, test.arr1, test.arr2, test.arr3, test.output)
				break
			}
		}
	}
}
