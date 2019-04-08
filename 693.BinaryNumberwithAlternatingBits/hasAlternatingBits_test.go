package main

import (
	"testing"
	"truman.com/leetcode/693.BinaryNumberwithAlternatingBits/hasAlternatingBits"
)

func TestHasAlternatingBits(t *testing.T) {
	tests := []struct{
		input int
		output bool
	}{
		{5,true},
		{7,false},
	}
	for _, test := range tests {
		ret := hasAlternatingBits.HasAlternatingBits(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}