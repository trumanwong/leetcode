package main

import (
	"testing"
	"truman.com/leetcode/860.LemonadeChange/lemonadeChange"
)

func TestLemonadeChange(t *testing.T)  {
	tests := []struct{
		input []int
		output bool
	} {
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5,5,10,10,20}, false},
	}

	for _, test := range tests {
		ret := lemonadeChange.LemonadeChange(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.input, test.output)
		}
	}
}