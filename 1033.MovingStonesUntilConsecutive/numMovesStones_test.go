package main

import (
	"testing"
	"truman.com/leetcode/1033.MovingStonesUntilConsecutive/numMovesStones"
)

func TestNumMovesStones(t *testing.T)  {
	tests := []struct{
		a int
		b int
		c int
		output []int
	} {
		{1,2,5,[]int{1,2}},
		{4,3,2,[]int{0,0}},
	}

	for _, test := range tests {
		ret := numMovesStones.NumMovesStones(test.a, test.b, test.c)
		for i, v := range ret {
			if test.output[i] != v {
				t.Errorf("Got %v for input (%d, %d, %d); expected %v", ret, test.a, test.b, test.c, test.output)
			}
		}
	}
}