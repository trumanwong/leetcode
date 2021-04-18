package trimMean

import "sort"

func TrimMean(arr []int) float64 {
	sort.Ints(arr)
	l := int(float64(len(arr)) * 0.05)
	r := len(arr) - l*2
	return float64(sum(arr[l:l+r])) / float64(r)
}

func sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}
