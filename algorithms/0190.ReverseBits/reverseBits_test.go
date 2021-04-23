package main

import (
	"leetcode/algorithms/0190.ReverseBits/reverseBits"
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		input  uint32
		output uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for _, test := range tests {
		ret := reverseBits.ReverseBits(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}
