package main

import (
	"testing"
	"truman.com/leetcode/1049.LastStoneWeightII/lastStoneWeightII"
)

func TestLastStoneWeightII(t *testing.T)  {
	tests := []struct{
		stones []int
		output int
	}{
		{[]int{2,7,4,1,8,1}, 1},
	}

	for _, test := range tests {
		ret := lastStoneWeightII.LastStoneWeightII(test.stones)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.stones, test.output)
		}
	}
}