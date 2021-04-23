package main

import (
	"leetcode/algorithms/1056.ConfusingNumber/confusingNumber"
	"testing"
)

func TestConfusingNumber(t *testing.T) {
	tests := []struct {
		N      int
		output bool
	}{
		{6, true},
		{89, true},
		{11, false},
		{25, false},
	}

	for _, test := range tests {
		ret := confusingNumber.ConfusingNumber(test.N)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.N, test.output)
		}
	}
}
