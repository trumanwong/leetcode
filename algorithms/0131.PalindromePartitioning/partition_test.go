package main

import (
	"leetcode/algorithms/0131.PalindromePartitioning/partition"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T)  {
	tests := []struct{
		s string
		output [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
	}

	for _, test := range tests {
		ret := partition.Partition(test.s)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %s; expected %v", ret, test.s, test.output)
		}
	}
}