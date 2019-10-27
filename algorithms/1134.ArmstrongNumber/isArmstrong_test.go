package _134_ArmstrongNumber

import (
	"leetcode/algorithms/1134.ArmstrongNumber/isArmstrong"
	"testing"
)

func TestIsArmstrong(t *testing.T)  {
	tests := []struct{
		input int
		output bool
	}{
		{153, true},
		{123, false},
	}

	for _, test := range tests {
		ret := isArmstrong.IsArmstrong(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %d; expected %t", ret, test.input, test.output)
		}
	}
}