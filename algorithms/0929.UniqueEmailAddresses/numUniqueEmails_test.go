package main

import (
	"leetcode/algorithms/0929.UniqueEmailAddresses/numUniqueEmails"
	"testing"
)

func TestNumUniqueEmails(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
	}

	for _, test := range tests {
		ret := numUniqueEmails.NumUniqueEmails(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}
