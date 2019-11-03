package maxLength

import (
	"sort"
	"strings"
)

func MaxLength(arr []string) int {
	newArr := make([]string, 0)
	for _, s := range arr {
		temp, judge := make(map[rune]int), false
		for _, v := range s {
			if _, ok := temp[v]; ok {
				judge = true
				break
			}
			temp[v]++
		}
		if !judge && s != "" {
			newArr = append(newArr, s)
		}
	}
	res := ""
	sort.Sort(MySort(newArr))
	for i := 0; i < len(newArr); i++ {
		temp := newArr[i]
		for j := 0; j < len(newArr) && i != j; j++ {
			if !strings.ContainsAny(temp, newArr[j]) {
				temp += newArr[j]
			}
		}

		if len(temp) > len(res) {
			res = temp
		}
	}
	return len(res)
}

type MySort []string

func (s MySort) Len() int {
	return len(s)
}
func (s MySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s MySort) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}