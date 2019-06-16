package main

import (
	"testing"
	"truman.com/leetcode/345.ReverseVowelsofaString/reverseVowels"
)

func TestReverseVowels(t *testing.T)  {
	tests := []struct{
		input string
		output string
	}{
		{"hello", "holle"},
		{"leetcode","leotcede"},
	}

	for _, test := range tests {
		ret := reverseVowels.ReverseVowels(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}