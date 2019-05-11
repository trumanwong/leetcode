package main

import (
	"testing"
	"truman.com/leetcode/1029.TwoCityScheduling/twoCitySchedCost"
)

func TestTwoCitySchedCost(t *testing.T)  {
	tests := []struct{
		input [][]int
		output int
	}{
		{[][]int{{259,770},{448,54},{926,667},{184,139},{840,118},{577,469}},1859},
	}

	for _, test := range tests {
		ret := twoCitySchedCost.TwoCitySchedCost(test.input)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.input, test.output)
		}
	}
}