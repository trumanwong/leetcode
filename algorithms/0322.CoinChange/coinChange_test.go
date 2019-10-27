package main

import (
	"leetcode/algorithms/0322.CoinChange/coinChange"
	"testing"
)

func TestCoinChange(t *testing.T)  {
	tests := []struct{
		coins []int
		amount int
		output int
	}{
		{[]int{1,2,5},11,3},
		{[]int{2},3,-1},
		{[]int{2,5,10,1},27,4},
		{[]int{186,419,83,408}, 6249,20},
	}
	for _, test := range tests {
		ret := coinChange.CoinChange(test.coins, test.amount)
		if ret != test.output {
			t.Errorf("Got %d for input %v, amount = %d; expected %d", ret, test.coins, test.amount, test.output)
		}
	}
}