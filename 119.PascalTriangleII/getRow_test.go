package main

import (
	"testing"
	"truman.com/leetcode/119.PascalTriangleII/getRow"
)

func TestGetRow(t *testing.T)  {
	tests := []struct{
		input int
		output []int
	}{
		{3,[]int{1,3,3,1}},
	}
	for _, test := range tests {
		res := getRow.GetRow(test.input)
		for i := 0; i < len(res); i++{
			if res[i] != test.output[i] {
				t.Errorf("Got %v for input %d; expected %v", res, test.input, test.output)
				break
			}
		}
	}
}