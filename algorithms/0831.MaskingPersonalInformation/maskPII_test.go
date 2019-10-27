package main

import (
	"leetcode/algorithms/0831.MaskingPersonalInformation/maskPII"
	"testing"
)

func TestMaskPII(t *testing.T)  {
	tests := []struct{
		S string
		output string
	}{
		{"LeetCode@LeetCode.com","l*****e@leetcode.com"},
		{"AB@qq.com", "a*****b@qq.com"},
		{"1(234)567-890", "***-***-7890"},
		{"86-(10)12345678", "+**-***-***-5678"},
	}

	for _, test := range tests {
		ret := maskPII.MaskPII(test.S)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.S, test.output)
		}
	}
}