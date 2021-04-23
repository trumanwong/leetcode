package main

import (
	"leetcode/algorithms/0599.MinimumIndexSumofTwoLists/findRestaurant"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	tests := []struct {
		list1  []string
		list2  []string
		output []string
	}{
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			[]string{"Shogun"}},
		{[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			[]string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}},
	}

	for _, test := range tests {
		ret := findRestaurant.FindRestaurant(test.list1, test.list2)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input list1 = %v, list2 = %v; expected %v", ret, test.list1, test.list2, test.output)
				break
			}
		}
	}
}
