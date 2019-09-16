package main

import (
	"testing"
	"truman.com/leetcode/605.CanPlaceFlowers/canPlaceFlowers"
)

func TestCanPlaceFlowers(t *testing.T)  {
	tests := []struct{
		flowerbed []int
		n int
		output bool
	}{
		{[]int{1,0,0,0,1}, 1, true},
		{[]int{1,0,0,0,1}, 2, false},
	}

	for _, test := range tests {
		ret := canPlaceFlowers.CanPlaceFlowers(test.flowerbed, test.n)
		if test.output != ret {
			t.Errorf("Got %t for input flowerbed = %v, n = %d; expected %t", ret, test.flowerbed, test.n, test.output)
		}
	}
}