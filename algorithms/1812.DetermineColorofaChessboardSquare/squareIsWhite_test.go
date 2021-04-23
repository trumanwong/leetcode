package main

import (
	"leetcode/algorithms/1812.DetermineColorofaChessboardSquare/squareIsWhite"
	"testing"
)

func TestSquareIsWhite(t *testing.T) {
	tests := []struct {
		coordinates string
		output      bool
	}{
		{"a1", false},
		{"b1", true},
		{"c7", false},
	}

	for _, test := range tests {
		ret := squareIsWhite.SquareIsWhite(test.coordinates)
		if ret != test.output {
			t.Errorf("Got %t for input %s; expected %t", ret, test.coordinates, test.output)
		}
	}
}
