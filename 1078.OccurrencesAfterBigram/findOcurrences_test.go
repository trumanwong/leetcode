package main

import (
	"fmt"
	"testing"
	"truman.com/leetcode/1078.OccurrencesAfterBigram/findOcurrences"
)

func TestFindOcurrences(t *testing.T)  {
	tests := []struct{
		text string
		first string
		second string
		output []string
	}{
		{"alice is a good girl she is a good student",  "a",  "good", []string{"girl","student"}},
		{"we will we will rock you",  "we",   "will", []string{"we","rock"}},
		{"alice is a good girl she is a good student",
			"a",
			"good", []string{"girl","student"}},
	}

	for _, test := range tests {
		ret := findOcurrences.FindOcurrences(test.text, test.first, test.second)
		fmt.Println(ret)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %s, first = %s, second = %s; expected %v", ret, test.text, test.first, test.second, test.output)
				break
			}
		}
	}
}