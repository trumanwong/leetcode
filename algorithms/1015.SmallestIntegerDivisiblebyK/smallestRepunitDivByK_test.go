package main

import (
	"leetcode/algorithms/1015.SmallestIntegerDivisiblebyK/smallestRepunitDivByK"
	"testing"
)

func TestSmallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		K      int
		output int
	}{
		{1, 1},
		{2, -1},
		{3, 3},
		{7, 6},
		{23, 22},
	}
	for _, test := range tests {
		ret := smallestRepunitDivByK.SmallestRepunitDivByK(test.K)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.K, test.output)
		}
	}
}
