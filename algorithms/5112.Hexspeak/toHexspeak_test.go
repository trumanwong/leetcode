package main

import (
	"leetcode/algorithms/5112.Hexspeak/toHexspeak"
	"testing"
)

func TestToHexspeak(t *testing.T)  {
	tests := []struct{
		num string
		output string
	}{
		{"257", "IOI"},
		{"14", "E"},
	}

	for _, test := range tests {
		ret := toHexspeak.ToHexspeak(test.num)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.num, test.output)
		}
	}
}