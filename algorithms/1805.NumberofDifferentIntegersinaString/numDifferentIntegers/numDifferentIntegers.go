package numDifferentIntegers

import (
	"strings"
)

func numDifferentIntegers(word string) int {
	res := make(map[string]int)
	s := make([]byte, 0)
	for i := 0; i < len(word); i++ {
		if word[i] >= '0' && word[i] <= '9' {
			s = append(s, word[i])
		} else {
			if len(s) > 0 {
				temp := strings.TrimLeft(string(s), "0")
				if len(temp) == 0 {
					res["0"]++
				} else {
					res[temp]++
				}
				s = make([]byte, 0)
			}
		}
	}
	if len(s) > 0 {
		temp := strings.TrimLeft(string(s), "0")
		if len(temp) == 0 {
			res["0"]++
		} else {
			res[temp]++
		}
	}
	return len(res)
}