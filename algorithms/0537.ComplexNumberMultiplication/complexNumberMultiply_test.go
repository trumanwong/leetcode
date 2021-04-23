package main

import (
	"leetcode/algorithms/0537.ComplexNumberMultiplication/complexNumberMultiply"
	"testing"
)

func TestComplexNumberMultiply(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
	}
	for _, test := range tests {
		ret := complexNumberMultiply.ComplexNumberMultiply(test.a, test.b)
		if ret != test.output {
			t.Errorf("Got %s for %s * %s; expected %s", ret, test.a, test.b, test.output)
		}
	}
}
