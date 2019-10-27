package main

import (
	"leetcode/algorithms/0985.SumofEvenNumbersAfterQueries/sumEvenAfterQueries"
	"testing"
)

func TestSumEvenAfterQueries(t *testing.T)  {
	tests := []struct{
		A []int
		queries [][]int
		output []int
	}{
		{[]int{1,2,3,4}, [][]int{{1,0},{-3,1},{-4,0},{2,3}}, []int{8,6,2,4}},
	}

	for _, test := range tests {
		ret := sumEvenAfterQueries.SumEvenAfterQueries(test.A, test.queries)
		for i, v := range ret {
			if test.output[i] != v {
				t.Errorf("Got %v for A=%v, queries=%v; expected %v", ret, test.A, test.queries, test.output)
				break
			}
		}
	}
}