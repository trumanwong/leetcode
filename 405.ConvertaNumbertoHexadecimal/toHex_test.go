package main

import (
	"testing"
	"truman.com/leetcode/405.ConvertaNumbertoHexadecimal/toHex"
)

func TestToHex(t *testing.T)  {
	tests := []struct{
		num int
		output string
	}{
		{26,"1a"},
		{-1,"ffffffff"},
	}

	for _, test := range tests {
		ret := toHex.ToHex(test.num)
		if ret != test.output {
			t.Errorf("Got %s for input %d; expected %s", ret, test.num, test.output)
		}
	}
}