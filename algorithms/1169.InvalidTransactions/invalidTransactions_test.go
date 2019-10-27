package main

import (
	"leetcode/algorithms/1169.InvalidTransactions/invalidTransactions"
	"testing"
)

func TestInvalidTransactions(t *testing.T)  {
	tests := []struct{
		transactions []string
		output []string
	}{
		{[]string{"alice,20,800,mtv","alice,50,100,beijing"},
			[]string{"alice,20,800,mtv","alice,50,100,beijing"}},
		{[]string{"alice,20,800,mtv","bob,50,1200,mtv"},
			[]string{"bob,50,1200,mtv"}},
	}

	for _, test := range tests {
		ret := invalidTransactions.InvalidTransactions(test.transactions)

		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %v; expected %v", ret, test.transactions, test.output)
				break
			}
		}
	}
}