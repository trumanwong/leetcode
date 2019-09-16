package main

import (
	"testing"
	"truman.com/leetcode/1108.DefanginganIPAddress/defangIPaddr"
)

func TestDefangIPaddr(t *testing.T)  {
	tests := []struct{
		address string
		output string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},

	}

	for _, test := range tests {
		ret := defangIPaddr.DefangIPaddr(test.address)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.address, test.output)
		}
	}
}