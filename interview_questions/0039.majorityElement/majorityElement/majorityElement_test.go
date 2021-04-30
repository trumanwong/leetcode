package majorityElement

import "testing"

func TestMajorityElement(t *testing.T)  {
	tests := []struct{
		nums []int
		output int
	}{
		{[]int{1, 2, 3, 2, 2, 2, 5, 4, 2}, 2},
	}

	for _, test := range tests {
		ret := majorityElement(test.nums)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.nums, test.output)
		}
	}
}