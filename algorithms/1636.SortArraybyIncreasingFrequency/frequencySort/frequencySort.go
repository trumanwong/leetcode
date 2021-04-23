package frequencySort

import "sort"

type NumFreq struct {
	num  int
	freq int
}

func FrequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	numFreqs := make([]NumFreq, 0)
	for k, v := range m {
		numFreqs = append(numFreqs, NumFreq{
			num:  k,
			freq: v,
		})
	}
	sort.Slice(numFreqs, func(i, j int) bool {
		if numFreqs[i].freq == numFreqs[j].freq {
			return numFreqs[i].num > numFreqs[j].num
		}
		return numFreqs[i].freq < numFreqs[j].freq
	})
	res := make([]int, 0)
	for _, numFreq := range numFreqs {
		for i := 0; i < numFreq.freq; i++ {
			res = append(res, numFreq.num)
		}
	}
	return res
}
