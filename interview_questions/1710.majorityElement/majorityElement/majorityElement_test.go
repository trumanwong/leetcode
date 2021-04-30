package majorityElement

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums   []int
		output int
	}{
		{[]int{1,2,5,9,5,9,5,5,5}, 5},
		{[]int{3,2},-1},
		{[]int{2,2,1,1,1,2,2}, 2},
		{[]int{3}, 3},
	}

	for _, test := range tests {
		ret := majorityElement(test.nums)
		if ret != test.output{
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}
