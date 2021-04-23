package main

import (
	"leetcode/algorithms/0165.CompareVersionNumbers/compareVersion"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	tests := []struct {
		version1 string
		version2 string
		output   int
	}{
		{"0.1", "1.1", -1},
		{"1.0.1", "1", 1},
		{"7.5.2.4", "7.5.3", -1},
		{"1.01", "1.001", 0},
		{"1.0", "1.0.0", 0},
	}

	for _, test := range tests {
		ret := compareVersion.CompareVersion(test.version1, test.version2)
		if ret != test.output {
			t.Errorf("Got %d for input version1 = %s, version = %s; expected %d", ret, test.version1, test.version2, test.output)
		}
	}
}
