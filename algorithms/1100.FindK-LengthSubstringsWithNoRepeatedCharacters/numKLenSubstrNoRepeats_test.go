package main

import (
	"leetcode/algorithms/1100.FindK-LengthSubstringsWithNoRepeatedCharacters/numKLenSubstrNoRepeats"
	"testing"
)

func TestNumKLenSubstrNoRepeats(t *testing.T) {
	tests := []struct {
		S      string
		K      int
		Output int
	}{
		{"havefunonleetcode", 5, 6},
		{"home", 5, 0},
	}

	for _, test := range tests {
		ret := numKLenSubstrNoRepeats.NumKLenSubstrNoRepeats(test.S, test.K)
		if ret != test.Output {
			t.Errorf("Got %d for input S = %s, K= %d; expected %d", ret, test.S, test.K, test.Output)
		}
	}
}
