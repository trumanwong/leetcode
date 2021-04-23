package main

import (
	"leetcode/algorithms/0093.RestoreIPAddresses/restoreIpAddresses"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		s      string
		output []string
	}{
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	}

	for _, test := range tests {
		ret := restoreIpAddresses.RestoreIpAddresses(test.s)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %s; expected %v", ret, test.s, test.output)
			}
		}
	}
}
