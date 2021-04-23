package main

import (
	"leetcode/algorithms/0043.MultiplyStrings/multiply"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		num1   string
		num2   string
		output string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"000", "0000", "0"},
		{"9", "9", "81"},
	}

	for _, test := range tests {
		ret := multiply.Multiply(test.num1, test.num2)
		if ret != test.output {
			t.Errorf("Got %s for input %s * %s; expected %s", ret, test.num1, test.num2, test.output)
		}
	}
}
