package main

import (
	"leetcode/algorithms/0086.PartitionList/partition"
	. "leetcode/common/list"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T)  {
	tests := []struct{
		head []int
		k int
		output []int
	}{
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
	}

	for _, test := range tests {
		ret := partition.Partition(Constructor(test.head), test.k)
		if !reflect.DeepEqual(ret.ToArray(), test.output) {
			t.Errorf("Got %v for input %v, k = %d; expected %v", ret, test.head, test.k, test.output)
		}
	}
}