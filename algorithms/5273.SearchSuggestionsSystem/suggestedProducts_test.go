package main

import (
	"leetcode/algorithms/5273.SearchSuggestionsSystem/suggestedProducts"
	"reflect"
	"testing"
)

func TestSuggestedProducts(t *testing.T)  {
	tests := []struct{
		products []string
		searchWord string
		output [][]string
	}{
		{[]string{"mobile","mouse","moneypot","monitor","mousepad"}, "mouse",
			[][]string{{"mobile","moneypot","monitor"},
			{"mobile","moneypot","monitor"},
			{"mouse","mousepad"},
			{"mouse","mousepad"},
			{"mouse","mousepad"}}},
			{[]string{"havana"}, "havana",
				[][]string{{"havana"},{"havana"},{"havana"},{"havana"},{"havana"},{"havana"}}},
			{[]string{"havana"}, "tatiana", [][]string{{},{},{},{},{},{},{}}},
			{[]string{"bags","baggage","banner","box","cloths"}, "bags",
				[][]string{{"baggage","bags","banner"},{"baggage","bags","banner"},{"baggage","bags"},{"bags"}}},
	}

	for _, test := range tests {
		ret := suggestedProducts.SuggestedProducts(test.products, test.searchWord)
		if !reflect.DeepEqual(ret, test.output) {
			t.Errorf("Got %v for input %v, searchWord = %s; expected %v", ret, test.products, test.searchWord, test.output)
		}
	}
}