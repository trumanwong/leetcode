package main

import (
	"leetcode/algorithms/1507.ReformatDate/reformatDate"
	"testing"
)

func TestReformatDate(t *testing.T)  {
	tests := []struct{
		input string
		output string
	}{
		{"20th Oct 2052", "2052-10-20"},
		{"6th Jun 1933", "1933-06-06"},
		{"26th May 1960", "1960-05-26"},
	}

	for _, test := range tests {
		ret := reformatDate.ReformatDate(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %v; expected %s", ret, test.input, test.output)
		}
	}
}