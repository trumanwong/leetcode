package groupAnagrams

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		temp := strings.Split(v, "")
		sort.Strings(temp)
		key := strings.Join(temp, "")
		m[key] = append(m[key], v)
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}