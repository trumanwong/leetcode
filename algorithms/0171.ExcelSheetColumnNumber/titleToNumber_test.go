package main

import (
	"leetcode/algorithms/0171.ExcelSheetColumnNumber/titleToNumber"
	"testing"
)

func TestTitleToNumber(t *testing.T)  {
	tests := []struct{
		input string
		output int
	}{
		{"A",1},
		{"AB",28},
		{"ZY",701},
	}

	for _, test := range tests {
		ret := titleToNumber.TitleToNumber(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.output)
		}
	}
}