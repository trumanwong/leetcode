package main

import (
	"testing"
	"truman.com/leetcode/1052.GrumpyBookstoreOwner/maxSatisfied"
)

func TestMaxSatisfied(t *testing.T)  {
	tests := []struct{
		customers []int
		grumpy []int
		X int
		output int
	} {
		{[]int{1,0,1,2,1,1,7,5},  []int{0,1,0,1,0,1,0,1}, 3, 16},
	}

	for _, test := range tests {
		ret := maxSatisfied.MaxSatisfied(test.customers, test.grumpy, test.X)
		if ret != test.output {
			t.Errorf("Got %d for customers: %v, grumpy: %v, X: %d; expected %d", ret, test.customers, test.grumpy, test.X, test.output)
		}
	}
}