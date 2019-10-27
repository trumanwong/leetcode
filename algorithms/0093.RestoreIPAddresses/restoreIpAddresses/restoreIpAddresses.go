package restoreIpAddresses

import "strings"

func RestoreIpAddresses(s string) []string {
	res := make([]string, 0)
	backtrack(make([]string, 0), s, &res)
	return res
}

func backtrack(putStr []string, s string, result *[]string) {
	if len(putStr) == 4 {
		if len(s) == 0 {
			*result = append(*result, strings.Join(putStr, "."))
		}
		return
	}

	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}

		str := s[0:i]
		if (len(str) == 3 && strings.Compare(str, "255") > 0) || (len(str) > 1 && s[0] == '0') {
			return
		}

		putStr = append(putStr, str)
		backtrack(putStr, s[i:], result)
		putStr = putStr[:len(putStr) - 1]
	}
}