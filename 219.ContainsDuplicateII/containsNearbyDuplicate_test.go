package main

import (
	"testing"
	"truman.com/leetcode/219.ContainsDuplicateII/containsNearbyDuplicate"
)

func TestContainsNearbyDuplicate(t *testing.T)  {
	tests := []struct{
		input []int
		K int
		outpu bool
	}{
		{[]int{1,2,3,1},3,true},
		{[]int{1,0,1,1},1,true},
		{[]int{1,2,3,1,2,3},2,false},
	}

	for _, test := range tests {
		ret := containsNearbyDuplicate.ContainsNearbyDuplicate(test.input, test.K)
		if ret != test.outpu {
			t.Errorf("Got %t for input %v, K = %d; expected %t", ret, test.input, test.K, test.outpu)
		}
	}
}