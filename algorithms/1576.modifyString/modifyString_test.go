package main

import (
	"leetcode/algorithms/1576.modifyString/modifyString"
	"testing"
)

func TestModifyString(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"?zs", "azs"},
		{"ubv?w", "ubvaw"},
		{"j?qg??b", "jaqgacb"},
		{"??yw?ipkj?", "abywaipkja"},
		{"?", "a"},
	}

	for _, test := range tests {
		ret := modifyString.ModifyString(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
