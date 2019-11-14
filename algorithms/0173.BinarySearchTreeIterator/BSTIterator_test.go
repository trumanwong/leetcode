package main

import (
	"leetcode/algorithms/0173.BinarySearchTreeIterator/BSTIterator"
	"leetcode/common/treeNode"
	"testing"
)

func TestBSTIterator(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]interface{}
		output []interface{}
	}{
		{[]string{"BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"},
			[][]interface{}{{7, 3, 15, nil, nil, 9, 20}, {nil}, {nil}, {nil}, {nil}, {nil}, {nil}, {nil}, {nil}, {nil}},
			[]interface{}{nil, 3, 7, true, 9, true, 15, true, 20, false},
		},
	}

	for _, test := range tests {
		iterator := BSTIterator.BSTIterator{}
		for i, operate := range test.operates {
			if operate == "BSTIterator" {
				iterator = BSTIterator.BSTIteratorConstructor(treeNode.Constructor(0, test.values[i]))
			} else if operate == "next" {
				ret := iterator.Next()
				if ret != test.output[i] {
					t.Errorf("Got %d for Next(); expected %d", ret, test.output[i])
				}
			} else if (operate == "hasNext") {
				ret := iterator.HasNext()
				if ret != test.output[i] {
					t.Errorf("Got %t for HasNext(); expected %t", ret, test.output[i])
				}
			}
		}
	}
}