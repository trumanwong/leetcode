package repeatedSubstringPattern

import "strings"

func RepeatedSubstringPattern(s string) bool {
	temp := s + s
	return strings.Contains(temp[1:len(temp) - 1], s)
}