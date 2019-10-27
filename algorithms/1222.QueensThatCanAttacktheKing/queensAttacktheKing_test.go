package main

import (
	"leetcode/algorithms/1222.QueensThatCanAttacktheKing/queensAttacktheKing"
	"testing"
)

func TestQueensAttacktheKing(t *testing.T)  {
	tests := []struct{
		queens [][]int
		king []int
		output [][]int
	}{
		{[][]int{{0,1},{1,0},{4,0},{0,4},{3,3},{2,4}}, []int{0, 0}, [][]int{{0,1},{1,0},{3,3}}},
		{[][]int{{0,0},{1,1},{2,2},{3,4},{3,5},{4,4},{4,5}}, []int{3, 3}, [][]int{{2,2},{3,4},{4,4}}},
	}

	for _, test := range tests {
		ret := queensAttacktheKing.QueensAttacktheKing(test.queens, test.king)
		judge := true
		for i, arr := range ret {
			for j, v := range arr {
				if v != test.output[i][j] {
					judge = false
					break
				}
			}
			if !judge {
				t.Errorf("Got %v for queens = %v, king = %v; expected %v", ret, test.queens, test.king, test.output)
			}
		}
	}
}