package main

import (
	"testing"
	"truman.com/leetcode/709.ToLowerCase/toLowerCase"
)

func TestToLowerCase(t *testing.T)  {
	tests := []struct{
		input string
		output string
	}{
		{"Hello","hello"},
		{"here","here"},
		{"LOVELY","lovely"},
	}

	for _, test := range tests {
		ret := toLowerCase.ToLowerCase(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}