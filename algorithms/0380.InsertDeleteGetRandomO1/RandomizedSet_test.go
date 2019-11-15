package main

import (
	"leetcode/algorithms/0380.InsertDeleteGetRandomO1/RandomizedSet"
	"testing"
)

func TestRandomizedSet(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]int
		output []interface{}
	}{
		{[]string{"RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"},
		[][]int{{},{1},{2},{2},{},{1},{2},{}}, []interface{}{nil,true,false,true,1,true,false,2}},
	}

	for _, test := range tests {
		set := RandomizedSet.RandomizedSet{}
		for i, operate := range test.operates {
			if operate == "RandomizedSet" {
				set = RandomizedSet.Constructor()
			} else if operate == "insert" {
				ret := set.Insert(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for insert(%d); expected %t", ret, test.values[i][0], test.output[i])
				}
			} else if operate == "remove" {
				ret := set.Remove(test.values[i][0])
				if ret != test.output[i] {
					t.Errorf("Got %t for remove(%d); expected %t", ret, test.values[i][0], test.output[i])
				}
			} else if operate == "getRandom" {
				ret := set.GetRandom()
				if ret != test.output[i] {
					t.Errorf("Got %d for getRandom; expected %d", ret, test.output[i])
				}
			}
		}
	}
}
