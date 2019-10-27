package main

import (
	"leetcode/algorithms/0088.MergeSortedArray/merge"
	"testing"
)

func TestMerge(t *testing.T)  {
	tests := []struct{
		nums1 []int
		m int
		nums2 []int
		n int
		output []int
	}{
		{[]int{1,2,3,0,0,0},3,[]int{2,5,6},3, []int{1,2,2,3,5,6}},
	}
	for _, test := range tests {
		merge.Merge(test.nums1, test.m, test.nums2, test.n)
		for i := 0; i < len(test.nums1); i++ {
			if test.nums1[i] != test.output[i] {
				t.Errorf("Got %v; expected %v", test.nums1, test.output)
			}
		}
	}
}
