package main

import (
	"leetcode/algorithms/1247.MinimumSwapstoMakeStringsEqual/minimumSwap"
	"testing"
)

func TestMinimumSwap(t *testing.T) {
	tests := []struct {
		s1     string
		s2     string
		output int
	}{
		{"xx", "yy", 1},
		{"xy", "yx", 2},
		{"xx", "xy", -1},
		{"xxyyxyxyxx", "xyyxyxxxyx", 4},
	}

	for _, test := range tests {
		ret := minimumSwap.MinimumSwap(test.s1, test.s2)
		if ret != test.output {
			t.Errorf("Got %d for input s1 = %s, s= %s; expected %d", ret, test.s1, test.s2, test.output)
		}
	}
}
