package main

import (
	"leetcode/algorithms/1155.NumberofDiceRollsWithTargetSum/numRollsToTarget"
	"testing"
)

func TestNumRollsToTarget(t *testing.T)  {
	tests := []struct{
		d int
		f int
		target int
		output int
	}{
		{1,6,3,1},
		{2,6,7,6},
		{2,5,10,1},
		{1,2,3,0},
		{30,30,500,222616187},
	}

	for _, test := range tests {
		ret := numRollsToTarget.NumRollsToTarget(test.d, test.f, test.target)
		if ret != test.output {
			t.Errorf("Got %d for d = %d, f = %d, target = %d; expected %d", ret, test.d, test.f, test.target, test.output)
		}
	}
}