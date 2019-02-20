package main

import (
	"testing"
	"truman.com/leetcode/countAndSay/countAndSay"
)

func TestCountAndSay(t *testing.T)  {
	tests := []struct{
		input int
		output string
	} {
		{1, "1"},
		{4,"1211"},
	}

	for _, test := range tests {
		ret := countAndSay.CountAndSay(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.input, test.output)
		}
	}
}