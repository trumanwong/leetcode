package main

import (
	"leetcode/algorithms/0763.PartitionLabels/partitionLabels"
	"testing"
)

func TestPartitionLabels(t *testing.T)  {
	tests := []struct{
		S string
		output []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}

	for _, test := range tests {
		ret := partitionLabels.PartitionLabels(test.S)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %s; expected %v", ret, test.S, test.output)
			}
		}
	}
}