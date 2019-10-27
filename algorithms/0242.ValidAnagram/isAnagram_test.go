package main

import (
	"leetcode/algorithms/0242.ValidAnagram/isAnagram"
	"testing"
)

func TestIsAnagram(t *testing.T)  {
	tests := []struct{
		s string
		t string
		output bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, test := range tests {
		ret := isAnagram.IsAnagram(test.s, test.t)
		if ret != test.output {
			t.Errorf("Got %t for s = %s, t = %s; expected %t", ret, test.s, test.t, test.output)
		}
	}
}