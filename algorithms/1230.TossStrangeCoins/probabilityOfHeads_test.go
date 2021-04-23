package main

import (
	"leetcode/algorithms/1230.TossStrangeCoins/probabilityOfHeads"
	"testing"
)

func TestProbabilityOfHeads(t *testing.T) {
	tests := []struct {
		prob   []float64
		target int
		output float64
	}{
		{[]float64{0.4}, 1, 0.4000},
		{[]float64{0.5, 0.5, 0.5, 0.5, 0.5}, 0, 0.03125},
	}

	for _, test := range tests {
		ret := probabilityOfHeads.ProbabilityOfHeads(test.prob, test.target)
		if ret != test.output {
			t.Errorf("Got %f for input %v, target = %d; expected %f", ret, test.prob, test.target, test.output)
		}
	}
}
