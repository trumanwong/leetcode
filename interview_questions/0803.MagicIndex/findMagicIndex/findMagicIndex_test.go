package findMagicIndex

import "testing"

func TestFindMagicIndex(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{0, 2, 3, 4, 5}, 0},
		{[]int{1, 1, 1}, 1},
	}

	for _, test := range tests {
		ret := findMagicIndex(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expecyted %d", ret, test.nums, test.output)
		}
	}
}
