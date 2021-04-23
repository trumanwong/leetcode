package removeSubfolders

import (
	"sort"
	"strings"
)

func RemoveSubfolders(folder []string) []string {
	sort.Sort(newString(folder))

	res, m := make([]string, 0), make(map[string]int)
	for _, v := range folder {
		arr := strings.Split(v, "/")
		temp, has := "", false
		for _, s := range arr {
			if s == "" {
				continue
			}
			temp += "/" + s
			if _, ok := m[temp]; ok {
				has = true
				break
			}
		}
		if !has {
			m[v]++
			res = append(res, v)
		}
	}
	return res
}

type newString []string

func (s newString) Len() int {
	return len(s)
}

func (s newString) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s newString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
