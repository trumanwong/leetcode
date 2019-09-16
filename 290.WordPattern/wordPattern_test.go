package main

import (
	"testing"
	"truman.com/leetcode/290.WordPattern/wordPattern"
)

func TestWordPattern(t *testing.T)  {
	tests := []struct{
		pattern string
		str string
		output bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}

	for _, test := range tests {
		ret := wordPattern.WordPattern(test.pattern, test.str)
		if ret != test.output {
			t.Errorf("Got %t for pattern = %s, str = %s; expected %t", ret, test.pattern, test.str, test.output)
		}
	}
}