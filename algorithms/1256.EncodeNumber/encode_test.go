package main

import (
	"leetcode/algorithms/1256.EncodeNumber/encode"
	"testing"
)

func TestEncode(t *testing.T)  {
	tests := []struct{
		num int
		output string
	}{
		{3, "00"},
		{23, "1000"},
		{107, "101100"},
		{235689047, "110000011000101010001011000"},
	}

	for _, test := range tests {
		ret := encode.Encode(test.num)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.num, test.output)
		}
	}
}