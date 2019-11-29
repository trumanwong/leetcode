package main

import (
	"leetcode/algorithms/0893.GroupsofSpecialEquivalentStrings/numSpecialEquivGroups"
	"testing"
)

func TestNumSpecialEquivGroups(t *testing.T)  {
	tests := []struct{
		A []string
		output int
	}{
		{[]string{"a","b","c","a","c","c"}, 3},
	}

	for _, test := range tests {
		ret := numSpecialEquivGroups.NumSpecialEquivGroups(test.A)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.A, test.output)
		}
	}
}