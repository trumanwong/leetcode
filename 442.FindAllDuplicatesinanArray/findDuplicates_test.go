package main

import (
	"testing"
	"truman.com/leetcode/442.FindAllDuplicatesinanArray/findDuplicates"
)

func TestFindDuplicates(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{4,3,2,7,8,2,3,1}, []int{2, 3}},
	}

	for _, test := range tests {
		ret := findDuplicates.FindDuplicates(test.input)
		if len(ret) != len(test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
		}
		for i := 0; i < len(ret); i++ {
			if ret[i] != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}