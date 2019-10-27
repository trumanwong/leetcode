package main

import (
	"leetcode/algorithms/0012.IntegertoRoman/intToRoman"
	"testing"
)

func TestIntToRoman(t *testing.T)  {
	tests := []struct{
		input int
		ans string
	} {
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{12, "XII"},
		{27, "XXVII"},
		{41, "XLI"},
		{58, "LVIII"},
		{190, "CXC"},
		{1994, "MCMXCIV"},
	}

	for _, tt := range tests {
		ret := intToRoman.IntToRoman(tt.input)
		if ret != tt.ans {
			t.Errorf("Got %s for input %d; expected %s", ret, tt.input, tt.ans)
		}
	}
}
