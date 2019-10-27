package main

import (
	"fmt"
	"leetcode/algorithms/1233.RemoveSubFoldersfromtheFilesystem/removeSubfolders"
	"testing"
)

func TestRemoveSubfolders(t *testing.T)  {
	tests := []struct{
		folder []string
		output []string
	}{
		{[]string{"/a","/a/b","/c/d","/c/d/e","/c/f"}, []string{"/a","/c/d","/c/f"}},
		{[]string{"/a","/a/b/c","/a/b/d"}, []string{"a"}},
		{[]string{"/a/b/c","/a/b/d","/a/b/ca"}, []string{"/a/b/c","/a/b/ca","/a/b/d"}},
	}

	for _, test := range tests {
		ret := removeSubfolders.RemoveSubfolders(test.folder)
		fmt.Println(ret)
	}
}