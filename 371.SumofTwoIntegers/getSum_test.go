package main

import (
	"testing"
	"truman.com/leetcode/371.SumofTwoIntegers/getSum"
)

func TestGetSum(t *testing.T)  {
	tests := []struct{
		a int
		b int
		output int
	}{
		{1,2,3},
		{-2,3,1},
	}

	for _, test := range tests {
		ret := getSum.GetSum(test.a, test.b)
		if ret != test.output {
			t.Errorf("Got %d for %d + %d; expected %d", ret, test.a, test.b, test.output)
		}
	}
}