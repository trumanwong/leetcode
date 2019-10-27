package main

import (
	"leetcode/algorithms/0443.StringCompression/compress"
	"testing"
)


func TestCompress(t *testing.T)  {
	tests := []struct{
		chars []byte
		output int
	}{
		{[]byte{'a','a','b','b','c','c','c'}, 6},
		{[]byte{'a'}, 1},
	}

	for _, test := range tests {
		ret := compress.Compress(test.chars)
		if ret != test.output {
			t.Errorf("Got %d for input %v; expected %d", ret, test.chars, test.output)
		}
	}
}