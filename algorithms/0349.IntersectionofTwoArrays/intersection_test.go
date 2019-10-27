package main

import (
	"leetcode/algorithms/0349.IntersectionofTwoArrays/intersection"
	"testing"
)

func TestIntersection(t *testing.T)  {
	tests := []struct{
		nums1 []int
		nums2 []int
		output []int
	}{
		{[]int{1,2,2,1},[]int{2,2},[]int{2}},
		{[]int{4,9,5},[]int{9,4,9,8,4},[]int{9,4}},
	}

	for _, test := range tests {
		ret := intersection.Intersection(test.nums1, test.nums2)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input (%v, %v); expected %v", ret, test.nums1, test.nums2, test.output)
				break
			}
		}
	}
}