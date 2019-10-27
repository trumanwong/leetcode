package main

import (
	"leetcode/algorithms/0063.UniquePathsII/uniquePathsWithObstacles"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T)  {
	tests := []struct{
		obstacleGrid [][]int
		output int
	}{
		{[][]int{{0,0,0}, {0,1,0}, {0,0,0}}, 2},
	}

	for _, test := range tests {
		ret := uniquePathsWithObstacles.UniquePathsWithObstacles(test.obstacleGrid)
		if ret != test.output {
			t.Errorf("Got %v for input %d; expected %d", ret, test.obstacleGrid, test.output)
		}
	}
}