package main

import (
	"testing"
	"truman.com/leetcode/6.ZigZagConversion/convert"
)

func TestConvert(t *testing.T)  {
	tests := []struct {
		s string
		numRows int
		output string
	}{
		{"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		{"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
		{"A", 2, "A"},
		{"AB", 3, "AB"},
	}

	for _, test := range tests {
		ret := convert.Convert(test.s, test.numRows)
		if ret != test.output {
			t.Errorf("Got %s for input: s = %s, numRows = %d; expected %s", ret, test.s, test.numRows, test.output)
		}
	}
}
