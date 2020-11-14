package main

import (
	"leetcode/algorithms/1598.CrawlerLogFolder/minOperations"
	"testing"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		logs   []string
		output int
	}{
		{[]string{"d1/", "d2/", "../", "d21/", "./"}, 2},
		{[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, 3},
		{[]string{"d1/", "../", "../", "../"}, 0},
	}

	for _, test := range tests {
		ret := minOperations.MinOperations(test.logs)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.logs, test.output)
		}
	}
}
