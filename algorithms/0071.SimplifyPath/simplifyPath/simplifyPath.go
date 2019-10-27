package simplifyPath

import "strings"

func SimplifyPath(path string) string {
	bytes := make([][]byte, 0)
	arr := strings.Split(path, "/")
	for _, v := range arr {
		if v != "" && v != "." && v != ".." {
			bytes = append(bytes, []byte(v))
		} else if v == ".." && len(bytes) > 0 {
			bytes = bytes[:len(bytes) - 1]
		}
	}

	res := "/"
	for i, v := range bytes {
		if i == len(bytes) - 1 {
			res += string(v)
			continue
		}
		res = res + string(v) + "/"
	}
	return res
}