package canMakeArithmeticProgression

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T)  {
	tests := []struct{
		arr []int
		output bool
	}{
		{[]int{3, 5, 1}, true},
		{[]int{1, 2, 4}, false},
	}

	for _, test := range tests {
		ret := CanMakeArithmeticProgression(test.arr)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.arr, test.output)
		}
	}
}