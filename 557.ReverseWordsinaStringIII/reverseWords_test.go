package main

import (
	"testing"
	"truman.com/leetcode/557.ReverseWordsinaStringIII/reverseWords"
)

func TestReverseWords(t *testing.T)  {
	tests := []struct{
		input string
		output string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, test := range tests {
		ret := reverseWords.ReverseWords(test.input)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.input, test.output)
		}
	}
}