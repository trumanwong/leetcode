package main

import (
	"testing"
	"truman.com/leetcode/389.FindtheDifference/findTheDifference"
)

func TestFindTheDifference(t *testing.T)  {
	tests := []struct{
		s string
		t string
		output string
	}{
		{"abcd", "abcde","e"},
	}
	for _, test := range tests {
		ret := findTheDifference.FindTheDifference(test.s, test.t)
		if string(ret) != test.output {
			t.Errorf("Got %s for s=%s,t=%s;expected %s", string(ret), test.s, test.t, test.output)
		}
	}
}