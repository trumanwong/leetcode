package makeGood

import (
	"strings"
)

func MakeGood(s string) string {
	res, b := []byte{}, []byte(s)
	for _, ch := range b {
		if len(res) > 0 && strings.ToLower(string(res[len(res) - 1])) == strings.ToLower(string(ch)) && res[len(res) - 1] != ch {
			res = res[:len(res) - 1]
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}