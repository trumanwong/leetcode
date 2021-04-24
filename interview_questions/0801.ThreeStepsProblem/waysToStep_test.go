package main

import (
	"leetcode/interview_questions/0801.ThreeStepsProblem/waysToStep"
	"testing"
)

func TestWaysToStep(t *testing.T)  {
	tests := []struct{
		n int
		output int
	}{
		{3, 4},
		{5, 13},
	}

	for _, test := range tests {
		ret := waysToStep.WaysToStep(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}