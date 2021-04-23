package groupThePeople

func GroupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	for i, v := range groupSizes {
		m[v] = append(m[v], i)
	}

	res := make([][]int, 0)
	for k, arr := range m {
		temp := make([]int, 0)
		for _, v := range arr {
			if len(temp) == k {
				res = append(res, temp)
				temp = make([]int, 0)
			}
			temp = append(temp, v)
		}
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}
	return res
}
