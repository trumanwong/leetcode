package main

import (
	"leetcode/algorithms/1154.DayoftheYear/ordinalOfDate"
	"testing"
)

func TestOrdinalOfDate(t *testing.T) {
	tests := []struct {
		date   string
		output int
	}{
		{"2019-01-09", 9},
		{"2019-02-10", 41},
		{"2003-03-01", 60},
		{"2004-03-01", 61},
	}

	for _, test := range tests {
		ret := ordinalOfDate.OrdinalOfDate(test.date)
		if ret != test.output {
			t.Errorf("Got %d for input %s; expected %d", ret, test.date, test.output)
		}
	}
}
