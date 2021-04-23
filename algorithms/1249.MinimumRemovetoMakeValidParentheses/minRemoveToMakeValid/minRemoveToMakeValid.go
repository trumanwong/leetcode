package minRemoveToMakeValid

import "fmt"

func MinRemoveToMakeValid(s string) string {
	l, r := 0, 0
	b := make([]rune, 0)
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			r++
			if r > l {
				r--
				continue
			}
		}
	}
	tl, tr := 0, 0
	equal := l
	if l > r {
		equal = r
	}
	fmt.Println(equal)
	for _, v := range s {
		if v == '(' {
			tl++
		} else if v == ')' {
			tr++
			if tr > tl {
				tr--
				continue
			}
		}

		if (v == '(' && tl > equal) || (v == ')' && tr > equal) {
			continue
		}
		b = append(b, v)
	}
	return string(b)
}
