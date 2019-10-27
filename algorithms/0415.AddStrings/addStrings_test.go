package main

import (
	"leetcode/algorithms/0415.AddStrings/addStrings"
	"testing"
)

func TestAddStrings(t *testing.T)  {
	tests := []struct{
		num1 string
		num2 string
		output string
	}{
		{"111","222","333"},
		{"191","919","1110"},
		{"19","101","120"},
	}
	for _, test := range tests {
		ret := addStrings.AddStrings(test.num1, test.num2)
		if ret != test.output {
			t.Errorf("Got %s for %s + %s; expected %s", ret, test.num1, test.num2, test.output)
		}
	}
}