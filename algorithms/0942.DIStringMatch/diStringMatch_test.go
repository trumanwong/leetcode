package main

import (
	"leetcode/algorithms/0942.DIStringMatch/diStringMatch"
	"testing"
)

func TestDiStringMatch(t *testing.T)  {
	tests := []struct {
		S string
		N []int
	}{
		{"IDID",[]int{0,4,1,3,2}},
		{"III",[]int{0,1,2,3}},
		{"DDI",[]int{3,2,0,1}},
	}

	for _, test := range tests {
		ret := diStringMatch.DiStringMatch(test.S)
		for i, v := range ret {
			if v != test.N[i] {
				t.Errorf("Got %v for input %s; expected %v", ret, test.S, test.N)
				break
			}
		}
	}
}