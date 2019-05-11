package main

import (
	"testing"
	"truman.com/leetcode/448.FindAllNumbersDisappearedinanArray/findDisappearedNumbers"
)

func TestFindDisappearedNumbers(t *testing.T)  {
	tests := []struct{
		input []int
		output []int
	}{
		{[]int{4,3,2,7,8,2,3,1},[]int{5,6}},
	}
	for _, test := range tests {
		ret := findDisappearedNumbers.FindDisappearedNumbers(test.input)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
			}
		}
	}
}