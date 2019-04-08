package main

import (
	"testing"
	"truman.com/leetcode/441.ArrangingCoins/arrangeCoins"
)

func TestArrangeCoins(t *testing.T)  {
	tests := []struct{
		input int
		output int
	}{
		{5,2},
		{8,3},
	}

	for _, test := range tests {
		ret := arrangeCoins.ArrangeCoins(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.input, test.output)
		}
	}
}