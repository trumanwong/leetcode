package main

import (
	"testing"
	"truman.com/leetcode/811.SubdomainVisitCount/subdomainVisits"
)

func TestSubdomainVisits(t *testing.T)  {
	tests := []struct{
		input []string
		output []string
	}{
		{[]string{"9001 discuss.leetcode.com"},
		[]string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}},
		{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
		[]string{"901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"}},
	}

	for _, test := range tests {
		ret := subdomainVisits.SubdomainVisits(test.input)
		if len(ret) != len(test.output) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
			continue
		}
		for _, v := range ret {
			if !inArray(v, test.output) {
				t.Errorf("Got %v for input %v; expected %v", ret, test.input, test.output)
				break
			}
		}
	}
}

func inArray(needle string, arr []string) bool {
	for _, v := range arr {
		if v == needle {
			return true
		}
	}
	return false
}