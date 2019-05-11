package main

import (
	"testing"
	"truman.com/leetcode/551.StudentAttendanceRecordI/checkRecord"
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