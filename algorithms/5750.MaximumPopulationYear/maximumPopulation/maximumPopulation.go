package maximumPopulation

func maximumPopulation(logs [][]int) int {
	m := make(map[int]int)
	year, max := 2051, 0
	for _, log := range logs {
		for i := log[0]; i < log[1]; i++ {
			m[i]++
			if m[i] >= max {
				if m[i] > max && i > year {
					year = i
				}
				max = m[i]
				if i < year {
					year = i
				}
			}
		}
	}
	return year
}