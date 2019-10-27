package main

import (
	"leetcode/algorithms/0925.LongPressedName/isLongPressedName"
	"testing"
)

func TestIsLongPressedName(t *testing.T)  {
	tests := []struct{
		name string
		typed string
		output bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
	}

	for _, test := range tests {
		ret := isLongPressedName.IsLongPressedName(test.name, test.typed)
		if ret != test.output {
			t.Errorf("Got %t for input name = %s, typed = %s; expected %t", ret, test.name, test.typed, test.output)
		}
	}
}