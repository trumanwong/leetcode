package main

import (
	"leetcode/algorithms/1047.RemoveAllAdjacentDuplicatesInString/removeDuplicates"
	"testing"
)

func TestRemoveDuplicates(t *testing.T)  {
	tests := []struct{
		S string
		output string
	} {
		{"abbaca", "ca"},
		{"aaaaaaaa", ""},
		{"aaaaaaaaa", "a"},
	}

	for _, test := range tests {
		ret := removeDuplicates.RemoveDuplicates(test.S)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.S, test.output)
		}
	}
}
