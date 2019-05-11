package main

import (
	"testing"
	"truman.com/leetcode/859.BuddyStrings/buddyStrings"
)

func TestBuddyStrings(t *testing.T)  {
	tests := []struct{
		A string
		B string
		output bool
	}{
		{"ab", "ba", true},
		{"ab", "ab", false},
		{"aaaaaaabc", "aaaaaaacb", true},
		{"","aa",false},
	}

	for _, test := range tests {
		ret := buddyStrings.BuddyStrings(test.A, test.B)
		if ret != test.output {
			t.Errorf("Got %t for A = %s, B = %s; expected %t", ret, test.A, test.B, test.output)
		}
	}
}