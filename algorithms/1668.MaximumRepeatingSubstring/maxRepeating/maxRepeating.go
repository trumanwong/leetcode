package maxRepeating

import "strings"

func maxRepeating(sequence string, word string) int {
	res := 0
	substr := word
	for strings.Contains(sequence, substr) {
		res++
		substr += word
	}
	return res
}