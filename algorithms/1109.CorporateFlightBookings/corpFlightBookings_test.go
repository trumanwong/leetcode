package main

import (
	"leetcode/algorithms/1109.CorporateFlightBookings/corpFlightBookings"
	"testing"
)

func TestCorpFlightBookings(t *testing.T)  {
	tests := []struct{
		bookings [][]int
		n int
		output []int
	}{
		{[][]int{{1,2,10},{2,3,20},{2,5,25}}, 5, []int{10,55,45,25,25}},
	}

	for _, test := range tests {
		ret := corpFlightBookings.CorpFlightBookings(test.bookings, test.n)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v, n = %d; expected %v", ret, test.bookings, test.n, test.output)
				break
			}
		}
	}
}