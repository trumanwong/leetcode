package main

import (
	"leetcode/algorithms/0709.ToLowerCase/toLowerCase"
	"testing"
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