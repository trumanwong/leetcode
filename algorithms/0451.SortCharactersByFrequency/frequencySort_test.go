package main

import (
	"leetcode/algorithms/0451.SortCharactersByFrequency/frequencySort"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		s      string
		output string
	}{
		{"tree", "eert"},
		{"cccaaa", "cccaaa"},
		{"Aabb", "bbaA"},
	}

	for _, test := range tests {
		ret := frequencySort.FrequencySort(test.s)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.s, test.output)
		}
	}
}
