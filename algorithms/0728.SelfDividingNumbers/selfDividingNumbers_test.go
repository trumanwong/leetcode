package main

import (
	"leetcode/algorithms/0728.SelfDividingNumbers/selfDividingNumbers"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	tests := []struct {
		left   int
		right  int
		output []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{66, 708, []int{66, 77, 88, 99, 111, 112, 115, 122, 124, 126, 128, 132, 135, 144, 155, 162, 168, 175, 184, 212, 216, 222, 224, 244, 248, 264, 288, 312, 315, 324, 333, 336, 366, 384, 396, 412, 424, 432, 444, 448, 488, 515, 555, 612, 624, 636, 648, 666, 672}},
	}

	for _, test := range tests {
		ret := selfDividingNumbers.SelfDividingNumbers(test.left, test.right)
		if len(ret) != len(test.output) {
			t.Errorf("Got %v for input (%d, %d); expected %v", ret, test.left, test.right, test.output)
		} else {
			for i := 0; i < len(ret); i++ {
				if ret[i] != test.output[i] {
					t.Errorf("Got %v for input (%d, %d); expected %v", ret, test.left, test.right, test.output)
					break
				}
			}
		}
	}
}
