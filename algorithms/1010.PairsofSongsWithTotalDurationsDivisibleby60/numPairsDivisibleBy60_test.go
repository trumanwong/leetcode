package main

import (
	"leetcode/algorithms/1010.PairsofSongsWithTotalDurationsDivisibleby60/numPairsDivisibleBy60"
	"testing"
)

func TestNumPairsDivisibleBy60(t *testing.T)  {
	tests := []struct{
		time []int
		output int
	}{
		{[]int{30,20,150,100,40}, 3},
		{[]int{60,60,60}, 3},
	}

	for _, test := range tests {
		ret := numPairsDivisibleBy60.NumPairsDivisibleBy60(test.time)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.time, test.output)
		}
	}
}