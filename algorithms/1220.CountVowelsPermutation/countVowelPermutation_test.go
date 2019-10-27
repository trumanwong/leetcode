package main

import (
	"leetcode/algorithms/1220.CountVowelsPermutation/countVowelPermutation"
	"testing"
)

func TestCountVowelPermutation(t *testing.T)  {
	tests := []struct{
		n int
		output int
	}{
		{1, 5},
		{2, 10},
	}

	for _, test := range tests {
		ret := countVowelPermutation.CountVowelPermutation(test.n)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.n, test.output)
		}
	}
}