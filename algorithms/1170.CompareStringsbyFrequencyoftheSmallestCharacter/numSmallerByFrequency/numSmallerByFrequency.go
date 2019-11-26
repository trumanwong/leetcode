package numSmallerByFrequency

import (
	"sort"
)

type ByteSlice []byte

func (b ByteSlice) Len() int {
	return len(b)
}

func (b ByteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func NumSmallerByFrequency(queries []string, words []string) []int {
	res := make([]int, 0)
	count := make([]int, 11)
	for i := 0; i < len(words); i++ {
		temp := ByteSlice([]byte(words[i]))
		sort.Sort(temp)
		count[f(temp)]++
	}

	for i := 9; i >= 0; i-- {
		count[i] += count[i + 1]
	}

	for _, v := range queries {
		temp := ByteSlice([]byte(v))
		sort.Sort(temp)
		res = append(res, count[f(temp) + 1])
	}
	return res
}

func f(s ByteSlice) int {
	res := 0
	sort.Sort(s)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i - 1] {
			break
		}
		res++
	}
	return res
}