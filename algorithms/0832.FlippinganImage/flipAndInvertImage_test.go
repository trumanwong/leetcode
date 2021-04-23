package main

import (
	"leetcode/algorithms/0832.FlippinganImage/flipAndInvertImage"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		A      [][]int
		output [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}

	for _, test := range tests {
		temp := make([][]int, len(test.A))
		for i, v := range test.A {
			temp[i] = make([]int, len(v))
			copy(temp[i], v)
		}

		ret := flipAndInvertImage.FlipAndInvertImage(test.A)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for input %v; expected %v", ret, test.A, test.output)
				break
			}
		}
	}
}
