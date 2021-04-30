package search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 6, 0},
	}

	for _, test := range tests {
		ret := search(test.nums, test.target)
		if ret != test.output {
			t.Errorf("Got %d for input (%v, %d); expected %d", ret, test.nums, test.target, test.output)
		}
	}
}
