package main

import (
	"testing"
	"truman.com/leetcode/1094.CarPooling/carPooling"
)

func TestCarPooling(t *testing.T)  {
	tests := []struct{
		trips [][]int
		capacity int
		output bool
	}{
		{[][]int{{2,1,5},{3,3,7}}, 4,false},
		{[][]int{{2,1,5},{3,3,7}}, 5,true},
		{[][]int{{2,1,5},{3,5,7}}, 3,true},
		{[][]int{{3,2,7},{3,7,8},{8,3,9}}, 11,true},
		{[][]int{{2,2,6},{2,4,7},{8,6,7}}, 11, true},
		{[][]int{{7,5,6},{6,7,8},{10,1,6}}, 16, false},
		{[][]int{{3,2,8},{4,4,6},{10,8,9}}, 11, true},
		{[][]int{{10,5,7},{10,3,4},{7,1,8},{6,3,4}}, 24, true},
	}

	for _, test := range tests {
		ret := carPooling.CarPooling(test.trips, test.capacity)
		if ret != test.output {
			t.Errorf("Got %t for trips: %v, capacity: %d; expected %t", ret, test.trips, test.capacity, test.output)
		}
	}
}