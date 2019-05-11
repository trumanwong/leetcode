package main

import (
	"testing"
	"truman.com/leetcode/804.UniqueMorseCodeWords/uniqueMorseRepresentations"
)

func TestUniqueMorseRepresentations(t *testing.T)  {
	tests := []struct{
		input []string
		output int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
	}

	for _, test := range tests {
		ret := uniqueMorseRepresentations.UniqueMorseRepresentations(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}