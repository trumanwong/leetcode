package main

import (
	"leetcode/algorithms/0914.XofaKindinaDeckofCards/hasGroupsSizeX"
	"testing"
)

func TestHasGroupsSizeX(t *testing.T) {
	tests := []struct {
		deck   []int
		output bool
	}{
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3}, false},
		{[]int{1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 1, 2, 2, 2, 2}, true},
	}

	for _, test := range tests {
		ret := hasGroupsSizeX.HasGroupsSizeX(test.deck)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.deck, test.output)
		}
	}
}
