package main

import (
	"leetcode/algorithms/1215.SteppingNumbers/countSteppingNumbers"
	"testing"
)

func TestCountSteppingNumbers(t *testing.T)  {
	tests := []struct{
		low int
		high int
		output []int
	}{
		{0,21,[]int{0,1,2,3,4,5,6,7,8,9,10,12,21}},
	}

	for _, test := range tests {
		res := countSteppingNumbers.CountSteppingNumbers(test.low, test.high)
		if len(res) != len(test.output) {
			t.Errorf("Got %v for input low: %d, high: %d, expected %v", res, test.low, test.high, test.output)
		}

		for i, v := range res {
			if test.output[i] != v {
				t.Errorf("Got %v for input low: %d, high: %d, expected %v", res, test.low, test.high, test.output)
				break
			}
		}
	}
}