package main

import (
	"fmt"
	"leetcode/common/list"
	"testing"
)

func TestList(t *testing.T)  {
	tests := []struct{
		nums []int
	}{
		{[]int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		ret := list.Constructor(test.nums)
		fmt.Println(ret)
	}
}