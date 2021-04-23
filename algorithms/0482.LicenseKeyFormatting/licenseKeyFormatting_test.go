package main

import (
	"leetcode/algorithms/0482.LicenseKeyFormatting/licenseKeyFormatting"
	"testing"
)

func TestLicenseKeyFormatting(t *testing.T) {
	tests := []struct {
		S      string
		K      int
		output string
	}{
		{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"2-5g-3-J", 2, "2-5G-3J"},
	}

	for _, test := range tests {
		ret := licenseKeyFormatting.LicenseKeyFormatting(test.S, test.K)
		if ret != test.output {
			t.Errorf("Got %s for S=%s, K=%d; expected %s", ret, test.S, test.K, test.output)
		}
	}
}
