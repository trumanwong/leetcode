package main

import (
	"leetcode/algorithms/0771.JewelsandStones/numJewelsInStones"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		J      string
		S      string
		output int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, test := range tests {
		ret := numJewelsInStones.NumJewelsInStones(test.J, test.S)
		if ret != test.output {
			t.Errorf("Got %d for input J = %s, S = %s;expected %d", ret, test.J, test.S, test.output)
		}
	}
}
