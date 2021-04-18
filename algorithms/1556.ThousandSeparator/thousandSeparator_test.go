package main

import (
	"leetcode/algorithms/1556.ThousandSeparator/thousandSeparator"
	"testing"
)

func TestThousandSeparator(t *testing.T)  {
	tests := []struct{
		input int
		output string
	}{
		{987, "987"},
		{1234, "1.234"},
		{123456789, "123.456.789"},
		{0, "0"},
		{1, "1"},
	}

	for _, test := range tests {
		ret := thousandSeparator.ThousandSeparator(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.input, test.output)
		}
	}
}