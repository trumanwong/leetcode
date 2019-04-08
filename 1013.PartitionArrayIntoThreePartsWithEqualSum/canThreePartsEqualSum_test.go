package main

import (
	"testing"
	"truman.com/leetcode/1013.PartitionArrayIntoThreePartsWithEqualSum/canThreePartsEqualSum"
)

func TestCanThreePartsEqualSum(t *testing.T)  {
	tests := []struct{
		A []int
		output bool
	}{
		{[]int{0,2,1,-6,6,-7,9,1,2,0,1}, true},
		{[]int{0,2,1,-6,6,7,9,-1,2,0,1}, false},
		{[]int{3,3,6,5,-2,2,5,1,-9,4}, true},
		{[]int{12,-4,16,-5,9,-3,3,8,0},true},
	}
	for _, test := range tests {
		ret := canThreePartsEqualSum.CanThreePartsEqualSum(test.A)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.A, test.output)
		}
	}
}