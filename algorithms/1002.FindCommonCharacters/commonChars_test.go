package main

import (
	"leetcode/algorithms/1002.FindCommonCharacters/commonChars"
	"testing"
)

func TestCommonChars(t *testing.T)  {
	tests := []struct{
		input []string
		output []string
	}{
		{[]string{"bella","label","roller"},[]string{"e","l","l"}},
		{[]string{"cool","lock","cook"},[]string{"c","o"}},
	}

	for _, test := range tests {
		ret := commonChars.CommonChars(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}