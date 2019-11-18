package main

import (
	"leetcode/algorithms/1261.FindElementsinaContaminatedBinaryTree/FindElements"
	"leetcode/common/treeNode"
	"testing"
)

func TestFindElements(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]interface{}
		output []interface{}
	}{
		{[]string{"FindElements","find","find"},
			[][]interface{}{{-1, nil, -1}, {1}, {2}},
			[]interface{}{nil, false, true}},
			{[]string{"FindElements","find","find","find"},
				[][]interface{}{{-1,-1,-1,-1,-1}, {1}, {3}, {5}},
				[]interface{}{nil, true, true, false}},
	}

	for _, test := range tests {
		elem := FindElements.FindElements{}
		for i, operate := range test.operates {
			if operate == "FindElements" {
				elem = FindElements.FindElementsConstructor(treeNode.Constructor(0, test.values[i]))
			} else if operate == "find" {
				ret := elem.Find(test.values[i][0].(int))
				if ret != test.output[i].(bool) {
					t.Errorf("Got %t for find(%d); expected %t", ret, test.values[i][0], test.output[i])
				}
			}
		}
	}
}