package main

import (
	"leetcode/algorithms/0067.AddBinary/addBinary"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for _, test := range tests {
		ret := addBinary.AddBinary(test.a, test.b)
		if ret != test.output {
			t.Errorf("Got %s for input %s + %s; expected %s", ret, test.a, test.b, test.output)
		}
	}
}
