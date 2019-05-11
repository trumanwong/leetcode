package main

import (
	"testing"
	"truman.com/leetcode/917.ReverseOnlyLetters/reverseOnlyLetters"
)

func TestReverseOnlyLetters(t *testing.T)  {
	tests := []struct{
		S string
		output string
	}{
		{"ab-cd", "dc-ba"},
		{ "a-bC-dEf-ghIj","j-Ih-gfE-dCba"},
		{"Test1ng-Leet=code-Q!","Qedo1ct-eeLg=ntse-T!"},
	}

	for _, test := range tests {
		ret := reverseOnlyLetters.ReverseOnlyLetters(test.S)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.S, test.output)
		}
	}
}