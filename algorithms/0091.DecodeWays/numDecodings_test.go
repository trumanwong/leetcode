package _091_DecodeWays

import (
	"leetcode/algorithms/0091.DecodeWays/numDecodings"
	"testing"
)

func TestNumDecodings(t *testing.T)  {
	tests := []struct{
		s string
		output int
	}{
		{"12", 2},
		{"226", 3},
	}

	for _, test := range tests {
		ret := numDecodings.NumDecodings(test.s)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.s, test.output)
		}
	}
}