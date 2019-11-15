package main

import (
	"leetcode/algorithms/0381.InsertDeleteGetRandomO1Duplicatesallowed/RandomizedCollection"
	"testing"
)

func TestRandomizedCollection(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]int
		output []interface{}
	}{
		{[]string{"RandomizedCollection","insert","insert","insert","getRandom","remove","getRandom"},
			[][]int{{},{1},{1},{2},{},{1},{}}, []interface{}{nil,true,false,true,1,true,2}},
	}

	for _, test := range tests {
		collection := RandomizedCollection.RandomizedCollection{}
		for i, operate := range test.operates {
			if operate == "RandomizedCollection" {
				collection = RandomizedCollection.Constructor()
			} else if operate == "insert" {
				ret := collection.Insert(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for insert(%d); expected %t", ret, test.values[i][0], test.output[i])
				}
			} else if operate == "remove" {
				ret := collection.Remove(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for remove(%d); expected %t", ret, test.values[i][0], test.output[i])
				}
			} else if operate == "getRandom" {
				ret := collection.GetRandom()
				if ret != test.output[i] {
					t.Errorf("Got %d for GetRandom(); expected %d", ret, test.output[i])
				}
			}
		}
	}
}