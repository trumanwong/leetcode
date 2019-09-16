package main

import (
	"testing"
	"truman.com/leetcode/1167.MinimumCosttoConnectSticks/connectSticks"
)

func TestConnectSticks(t *testing.T)  {
	tests := []struct{
		sticks []int
		output int
	}{
		{[]int{2, 4, 3}, 14},
		{[]int{1,8,3,5}, 30},
	}

	for _, test := range tests {
		ret := connectSticks.ConnectSticks(test.sticks)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.sticks, test.output)
		}
	}
}