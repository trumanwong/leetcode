package main

import (
	"leetcode/algorithms/0383.RansomNote/canConstruct"
	"testing"
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