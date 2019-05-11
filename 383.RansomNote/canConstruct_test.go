package main

import (
	"testing"
	"truman.com/leetcode/383.RansomNote/canConstruct"
)

func TestCanConstruct(t *testing.T)  {
	tests := []struct{
		ransomNote string
		magazine string
		output bool
	}{
		{"a", "b",false},
		{"aa", "ab",false},
		{"aa", "aab",true},
	}

	for _, test := range tests {
		ret := canConstruct.CanConstruct(test.ransomNote, test.magazine)
		if ret != test.output {
			t.Errorf("Got %t for ransomNote=%s, magazine=%s; expected %t", ret, test.ransomNote, test.magazine, test.output)
		}
	}
}