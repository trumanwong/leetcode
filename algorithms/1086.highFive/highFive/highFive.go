package highFive

import "sort"

func HighFive(items [][]int) [][]int {
	m := make(map[int][]int)
	ids := []int{}
	for i := 0; i < len(items); i++ {
		key := items[i][0]
		if _, ok := m[key]; !ok {
			ids = append(ids, items[i][0])
		}
		m[key] = append(m[key], items[i][1])
	}

	sort.Ints(ids)
	for k, _ := range m {
		sort.Ints(m[k])
	}

	res := [][]int{}
	for i := 0; i < len(ids); i++ {
		sum, size := 0, len(m[ids[i]])
		for j := size - 1; j >= size-5; j-- {
			sum += m[ids[i]][j]
		}
		temp := []int{ids[i], sum / 5}
		res = append(res, temp)
	}
	return res
}
