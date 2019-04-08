package main

import (
	"testing"
	"truman.com/leetcode/392.IsSubsequence/isSubsequence"
)

func TestIsSubsequence(t *testing.T)  {
	tests := []struct{
		s string
		t string
		output bool
	}{
		{"abc","ahbgdc",true},
		{"axc","ahbgdc",false},
	}

	for _, test := range tests {
		ret := isSubsequence.IsSubsequence(test.s, test.t)
		if ret != test.output {
			t.Errorf("Got %t for s = %s, t = %s; expected %t", ret, test.s, test.t, test.output)
		}
	}
}
