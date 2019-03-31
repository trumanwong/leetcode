package main

import (
	"testing"
	"truman.com/leetcode/8.StringtoInteger/myAtoi"
)

func TestMyAtoi(t *testing.T) {
	tests := [] struct {
		input string
		ans int
	} {
		{"", 0},
		{"   Hello World", 0},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"20000000000000000000", 2147483647 },
		{"  0000000000012345678", 12345678},
	}

	for _, tt := range tests {
		ret := myAtoi.MyAtoi(tt.input)
		if ret != tt.ans {
			t.Errorf("Got %d for input %s; expected %d", ret, tt.input, tt.ans)
		}
	}
}