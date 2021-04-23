package main

import (
	"leetcode/algorithms/0937.ReorderDatainLogFiles/reorderLogFiles"
	"testing"
)

func TestReorderLogFiles(t *testing.T) {
	tests := []struct {
		logs   []string
		output []string
	}{
		{[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			[]string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}},
	}

	for _, test := range tests {
		ret := reorderLogFiles.ReorderLogFiles(test.logs)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.logs, test.output)
				break
			}
		}
	}
}
