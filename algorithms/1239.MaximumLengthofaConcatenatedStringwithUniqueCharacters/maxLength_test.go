package main

import (
	"leetcode/algorithms/1239.MaximumLengthofaConcatenatedStringwithUniqueCharacters/maxLength"
	"testing"
)

func TestMaxLength(t *testing.T)  {
	tests := []struct{
		arr []string
		output int
	}{
		{[]string{"un","iq","ue"}, 4},
		{[]string{"cha","r","act","ers"}, 6},
		{[]string{"abcdefghijklmnopqrstuvwxyz"}, 26},
	}

	for _, test := range tests {
		ret := maxLength.MaxLength(test.arr)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.arr, test.output)
		}
	}
}