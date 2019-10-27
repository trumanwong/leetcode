package findRestaurant

func FindRestaurant(list1 []string, list2 []string) []string {
	m, m1, m2 := make(map[string]int), make(map[string]int), make(map[string]int)
	for i, v := range list1 {
		m[v] = i
		m1[v] = i
	}

	min := 100000
	for i, v := range list2 {
		if _, ok := m[v]; !ok {
			continue
		}
		m[v] = m[v] + i
		m2[v] = i

		if min > m[v] {
			min = m[v]
		}
	}

	ret := []string{}
	for i, v := range m {
		_, ok1 := m1[i]
		_, ok2 := m2[i]
		if !ok1 || !ok2 {
			continue
		}
		if v == min {
			ret = append(ret, i)
		}
	}
	return ret
}