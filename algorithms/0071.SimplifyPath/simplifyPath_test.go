package main

import (
	"leetcode/algorithms/0071.SimplifyPath/simplifyPath"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path   string
		output string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
	}

	for _, test := range tests {
		ret := simplifyPath.SimplifyPath(test.path)
		if ret != test.output {
			t.Errorf("Got %s for input %s; expected %s", ret, test.path, test.output)
		}
	}
}
