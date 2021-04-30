package minArray

import "testing"

func TestMinArray(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
	}

	for _, test := range tests {
		ret := minArray(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}
