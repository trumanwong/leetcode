package main

import (
	"leetcode/algorithms/1118.NumberofDaysinaMonth/numberOfDays"
	"testing"
)

func TestNumberOfDays(t *testing.T) {
	tests := []struct {
		Y      int
		M      int
		output int
	}{
		{2019, 7, 31},
	}

	for _, test := range tests {
		ret := numberOfDays.NumberOfDays(test.Y, test.M)
		if ret != test.output {
			t.Errorf("Got %d for Y = %d, M = %d;expected %d", ret, test.Y, test.M, test.output)
		}
	}
}
