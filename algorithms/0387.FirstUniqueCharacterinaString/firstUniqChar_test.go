package main

import (
	"leetcode/algorithms/0387.FirstUniqueCharacterinaString/firstUniqChar"
	"testing"
)

func TestFirstUniqChar(t *testing.T)  {
	tests := []struct{
		input string
		output int
	}{
		{"leetcode",0},
		{"loveleetcode",2},
	}

	for _, test := range tests {
		ret := firstUniqChar.FirstUniqChar(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.input, test.output)
		}
	}
}