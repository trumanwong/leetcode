package main

import (
	"leetcode/algorithms/0405.ConvertaNumbertoHexadecimal/toHex"
	"testing"
)

func TestToHex(t *testing.T) {
	tests := []struct {
		num    int
		output string
	}{
		{26, "1a"},
		{-1, "ffffffff"},
	}

	for _, test := range tests {
		ret := toHex.ToHex(test.num)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.num, test.output)
		}
	}
}
