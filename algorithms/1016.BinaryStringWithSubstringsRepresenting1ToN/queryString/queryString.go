package queryString

import (
	"strconv"
	"strings"
)

func QueryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		temp := strconv.FormatInt(int64(i),2)
		if !strings.Contains(S, temp) {
			return false
		}
	}
	return true
}