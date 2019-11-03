package main

import (
	"leetcode/algorithms/5096.ArrayTransformation/transformArray"
	"reflect"
	"testing"
)

func TestTransformArray(t *testing.T)  {
	tests := []struct{
		arr []int
		output []int
	}{
		{[]int{6,2, 3, 4}, []int{6, 3, 3, 4}},
		{[]int{1, 6, 3 , 4, 3, 5}, []int{1,4,4,4,4,5}},
		{[]int{2,1,2,1,1,2,2,1}, []int{2,2,1,1,1,2,2,1}},
	}

	for _, test := range tests {
		ret := transformArray.TransformArray(test.arr)
		if !reflect.DeepEqual(test.output, ret) {
			t.Errorf("Got %v for input %v; expected %v", ret, test.arr, test.output)
		}
	}
}