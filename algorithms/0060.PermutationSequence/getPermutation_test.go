package main

import (
	"leetcode/algorithms/0060.PermutationSequence/getPermutation"
	"testing"
)

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		output string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
	}

	for _, test := range tests {
		ret := getPermutation.GetPermutation(test.n, test.k)
		if ret != test.output {
			t.Errorf("Got %s for n = %d, k = %d; expected %s", ret, test.n, test.k, test.output)
		}
	}
}
