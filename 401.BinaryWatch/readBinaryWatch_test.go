package main

import (
	"testing"
	"truman.com/leetcode/401.BinaryWatch/readBinaryWatch"
)

func TestReadBinaryWatch(t *testing.T)  {
	tests := []struct{
		num int
		output []string
	}{
		{1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
	}

	for _, test := range tests {
		ret := readBinaryWatch.ReadBinaryWatch(test.num)
		for i, v := range ret {
			if v != test.output[i] {
				t.Errorf("Got %v for input %d; expected %v", ret, test.num, test.output)
				break
			}
		}
	}
}