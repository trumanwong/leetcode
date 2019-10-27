package main

import (
	"leetcode/algorithms/0551.StudentAttendanceRecordI/checkRecord"
	"testing"
)

func TestCheckRecord(t *testing.T)  {
	tests := []struct{
		input string
		output bool
	}{
		{"PPALLP", true},
		{"PPALLL", false},
	}

	for _, test := range tests {
		ret := checkRecord.CheckRecord(test.input)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.input, test.output)
		}
	}
}