package main

import (
	"leetcode/algorithms/1101.TheEarliestMomentWhenEveryoneBecomeFriends/earliestAcq"
	"testing"
)

func TestEarliestAcq(t *testing.T) {
	tests := []struct {
		logs   [][]int
		N      int
		Output int
	}{
		{[][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5},
			{20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6,
			20190301},
	}

	for _, test := range tests {
		ret := earliestAcq.EarliestAcq(test.logs, test.N)
		if ret != test.Output {
			t.Errorf("Got %d for input logs = %v, N = %d; expected %d", ret, test.logs, test.N, test.Output)
		}
	}
}
