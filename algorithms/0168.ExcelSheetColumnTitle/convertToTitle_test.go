package main

import (
	"leetcode/algorithms/0168.ExcelSheetColumnTitle/convertToTitle"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		input  int
		output string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
	}

	for _, test := range tests {
		ret := convertToTitle.ConvertToTitle(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.input, test.output)
		}
	}
}
