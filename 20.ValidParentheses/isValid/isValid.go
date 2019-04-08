package isValid

import "strings"

func IsValid(s string) bool {
	for strings.Index(s, "()") >= 0 || strings.Index(s, "[]") >= 0 || strings.Index(s, "{}") >= 0 {
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
	}
	return s == ""
}