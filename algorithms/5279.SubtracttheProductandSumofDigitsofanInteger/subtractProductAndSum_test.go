package main

import (
	"leetcode/algorithms/5279.SubtracttheProductandSumofDigitsofanInteger/subtractProductAndSum"
	"testing"
)

func TestSubtractProductAndSum(t *testing.T) {
	tests := []struct {
		n      int
		output int
	}{
		{234, 15},
	}

	for _, test := range tests {
		ret := subtractProductAndSum.SubtractProductAndSum(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}
