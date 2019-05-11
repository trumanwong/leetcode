package main

import (
	"testing"
	"truman.com/leetcode/796.RotateString/rotateString"
)

func TestRotateString(t *testing.T)  {
	tests := []struct{
		A string
		B string
		output bool
	}{
		{"abcde", "cdeab",true},
		{"abcde", "abced",false},
	}

	for _, test := range tests {
		ret := rotateString.RotateString(test.A, test.B)
		if ret != test.output {
			t.Errorf("Got %t for A=%s, B=%s; expected %t", ret, test.A, test.B, test.output)
		}
	}
}