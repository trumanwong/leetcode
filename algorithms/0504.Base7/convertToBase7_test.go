package main

import (
	"leetcode/algorithms/0504.Base7/convertToBase7"
	"testing"
)

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{100, "202"},
		{-7, "-10"},
	}

	for _, test := range tests {
		ret := convertToBase7.ConvertToBase7(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.input, test.output)
		}
	}
}
