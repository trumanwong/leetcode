package main

import (
	. "leetcode/common/treeNode"
	"reflect"
	"testing"
)

func TestPreOrderTraversal(t *testing.T)  {
	tests := []struct{
		nums []interface{}
		output []int
	}{
		{[]interface{}{5,3,6,2,4,nil,7}, []int{5, 3, 2, 4, 6, 7}},
	}

	for _, test := range tests {
		root := Constructor(0, test.nums)
		ret := root.PreOrderTraversal()
		if !reflect.DeepEqual(test.output, ret) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
		}
	}
}

func TestInOrderTraversal(t *testing.T)  {
	tests := []struct{
		nums []interface{}
		output []int
	}{
		{[]interface{}{5,3,6,2,4,nil,7}, []int{2, 3, 4, 5, 6, 7}},
	}

	for _, test := range tests {
		root := Constructor(0, test.nums)
		ret := root.InOrderTraversal()
		if !reflect.DeepEqual(test.output, ret) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
		}
	}
}

func TestLevelOrder(t *testing.T)  {
	tests := []struct{
		nums []interface{}
		output [][]int
	}{
		{[]interface{}{5,3,6,2,4,nil,7}, [][]int{{5}, {3, 6}, {2, 4, 7}}},
	}

	for _, test := range tests {
		root := Constructor(0, test.nums)
		ret := root.LevelOrder()
		if !reflect.DeepEqual(test.output, ret) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
		}
	}
}

func TestPostOrderTraversal(t *testing.T)  {
	tests := []struct{
		nums []interface{}
		output []int
	}{
		{[]interface{}{5,3,6,2,4,nil,7}, []int{2, 4, 3, 7, 6, 5}},
	}

	for _, test := range tests {
		root := Constructor(0, test.nums)
		ret := root.PostOrderTraversal()
		if !reflect.DeepEqual(test.output, ret) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.nums, test.output)
		}
	}
}