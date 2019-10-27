package main

import (
	"leetcode/algorithms/0944.DeleteColumnstoMakeSorted/minDeletionSize"
	"testing"
)

func TestMinDeletionSize(t *testing.T)  {
	tests := []struct{
		A []string
		output int
	}{
		{[]string{"cba", "daf", "ghi"}, 1},
		{[]string{"a", "b"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 3},
	}

	for _, test := range tests {
		ret := minDeletionSize.MinDeletionSize(test.A)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.A, test.output)
		}
	}
}