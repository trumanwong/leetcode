package main

import (
	"leetcode/algorithms/0507.PerfectNumber/checkPerfectNumber"
	"testing"
)

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		num    int
		output bool
	}{
		{28, true},
	}

	for _, test := range tests {
		ret := checkPerfectNumber.CheckPerfectNumber(test.num)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.num, test.output)
		}
	}
}
