package main

import (
	"testing"
	"truman.com/leetcode/884.UncommonWordsfromTwoSentences/uncommonFromSentences"
)

func TestUncommonFromSentences(t *testing.T)  {
	tests := []struct{
		A string
		B string
		output []string
	}{
		{"this apple is sweet", "this apple is sour", []string{"sweet","sour"}},
		{"apple apple",  "banana", []string{"banana"}},
	}

	for _, test := range tests {
		ret := uncommonFromSentences.UncommonFromSentences(test.A, test.B)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for A=%s, B=%s; expected %v", ret, test.A, test.B, test.output)
				break
			}
		}
	}
}