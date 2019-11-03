package main

import (
	"leetcode/algorithms/0234.PalindromeLinkedList/isPalindrome"
	. "leetcode/common/list"
	"testing"
)

func TestIsPalindrome(t *testing.T)  {
	tests := []struct{
		head []int
		output bool
	}{
		{[]int{1, 2}, false},
		{[]int{1, 2, 2, 1}, true},
	}

	for _, test := range tests {
		ret := isPalindrome.IsPalindrome(Constructor(test.head))
		if ret != test.output {
			t.Errorf("Got %t for input %v; expected %t", ret, test.head, test.output)
		}
	}
}