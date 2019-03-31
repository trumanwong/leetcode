package _84StringWithoutAAAorBBB

import (
	"testing"
	"truman.com/leetcode/984StringWithoutAAAorBBB/strWithout3a3b"
)

func TestStrWithout3a3b(t *testing.T)  {
	tests := []struct{
		A int
		B int
		output string
	}{
		{1,2,"bab"},
		{4,1,"aabaa"},
	}

	for _, test := range tests {
		ret := strWithout3a3b.StrWithout3a3b(test.A, test.B)
		if ret != test.output {
			t.Errorf("Got %s; expected %s", ret, test.output)
		}
	}
}
