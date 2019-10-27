package main

import (
	"leetcode/algorithms/0220.ContainsDuplicateIII/containsNearbyAlmostDuplicate"
	"testing"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T)  {
	tests := []struct{
		nums []int
		k int
		t int
		output bool
	}{
		{[]int{1,2,3,1},3,0,true},
		{[]int{1,0,1,1},1,2,true},
		{[]int{1,5,9,1,5,9},2,3,false},
	}

	for _, test := range tests {
		ret := containsNearbyAlmostDuplicate.ContainsNearbyAlmostDuplicate(test.nums, test.k, test.t)
		if ret != test.output {
			t.Errorf("Got %t for input nums = %v, k = %d, t = %d; expected %t", ret, test.nums, test.k, test.t, test.output)
		}
	}
}