package main

import (
	"testing"
	"truman.com/leetcode/1016.BinaryStringWithSubstringsRepresenting1ToN/queryString"
)

func TestQueryString(t *testing.T)  {
	tests := []struct{
		S string
		N int
		output bool
	}{
		{"0110",3,true},
		{"0110",4,false},
	}
	for _, test := range tests {
		ret := queryString.QueryString(test.S, test.N)
		if ret != test.output {
			t.Errorf("Got %t for input S = %s, N = %d; expected %t", ret, test.S, test.N, test.output)
		}
	}
}