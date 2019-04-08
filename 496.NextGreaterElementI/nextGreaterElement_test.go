package main

import (
	"testing"
	"truman.com/leetcode/496.NextGreaterElementI/nextGreaterElement"
)

func TestNextGreaterElement(t *testing.T)  {
	tests := []struct{
		nums1 []int
		nums2 []int
		output []int
	}{
		{[]int{4,1,2},[]int{1,3,4,2},[]int{-1,3,-1}},
	}

	for _, test := range tests {
		ret := nextGreaterElement.NextGreaterElement(test.nums1, test.nums2)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input nums1 = %v, nums2 = %v; expected %v", ret, test.nums1, test.nums2, test.output)
				break
			}
		}
	}
}
