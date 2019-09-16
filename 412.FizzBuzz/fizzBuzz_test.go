package main

import (
	"testing"
	"truman.com/leetcode/412.FizzBuzz/fizzBuzz"
)

func TestFizzBuzz(t *testing.T)  {
	tests := []struct{
		n int
		output []string
	}{
		{15, []string{"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz",}},
	}

	for _, test := range tests {
		ret := fizzBuzz.FizzBuzz(test.n)
		for i, v := range ret {
			if test.output[i] != v {
				t.Errorf("Got %v for input %d; expected %v", ret, test.n, test.output)
				break
			}
		}
	}
}