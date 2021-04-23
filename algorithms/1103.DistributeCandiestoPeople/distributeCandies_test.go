package main

import (
	"leetcode/algorithms/1103.DistributeCandiestoPeople/distributeCandies"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		candies    int
		num_people int
		output     []int
	}{
		{7, 4, []int{1, 2, 3, 1}},
		{10, 3, []int{5, 2, 3}},
	}

	for _, test := range tests {
		ret := distributeCandies.DistributeCandies(test.candies, test.num_people)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input candies = %d, num_people = %d; expected %v", ret, test.candies, test.num_people, test.output)
				break
			}
		}
	}
}
