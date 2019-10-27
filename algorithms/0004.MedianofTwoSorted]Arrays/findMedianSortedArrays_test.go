package main

import (
	"leetcode/algorithms/0004.MedianofTwoSorted]Arrays/findMedianSortedArrays"
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T)  {
	tests := []struct{
		nums1 []int
		nums2 []int
		output float64
	} {
		{[]int{1,3}, []int{2}, 2},
	}

	for _, test := range tests {
		ret := findMedianSortedArrays.FindMedianSortedArrays(test.nums1, test.nums2)
		if math.Abs(ret - test.output) > 0.0001 {
			t.Errorf("Got %d for input nums1 = %v, nums2 = %v; expected %d", ret, test.nums1, test.nums2, test.output)
		}
	}
}