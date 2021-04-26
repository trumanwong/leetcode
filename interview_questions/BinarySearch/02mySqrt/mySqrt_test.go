package main

import (
	"leetcode/interview_questions/BinarySearch/02mySqrt/mySqrt"
	"testing"
)

func TestMySqrt(t *testing.T)  {
	tests := []struct{
		x int
		output int
	}{
		{4, 2},
		{8, 2},
	}

	for _, test := range tests {
		ret := mySqrt.MySqrt(test.x)
		if ret != test.output {
			t.Errorf("Got %d for input %d; expected %d", ret, test.x, test.output)
		}
	}
}