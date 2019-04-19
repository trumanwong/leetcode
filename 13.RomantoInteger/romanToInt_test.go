package main

import (
	"testing"
	"truman.com/leetcode/13.RomantoInteger/romanToInt"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input string
		result int
	}{
		{"I", 1},
		{"II", 2},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"CM", 900},
	}

	for _, test := range tests {
		ret := romanToInt.RomanToInt(test.input)
		if ret != test.result {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.result)
		}
	}
}