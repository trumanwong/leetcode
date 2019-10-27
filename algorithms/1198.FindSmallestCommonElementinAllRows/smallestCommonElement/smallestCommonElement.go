package smallestCommonElement

func smallestCommonElement(mat [][]int) int {
	m := make(map[int]int)
	ret := 1000000
	for i := 0; i < len(mat); i++ {
		temp := make(map[int]int)
		for j := 0; j < len(mat[i]); j++ {
			v := mat[i][j]
			temp[v]++
			if i == 0 {
				m[v]++
			}
		}
		for k, _ := range m {
			if _, ok := temp[k]; !ok {
				delete(m, k)
			}
		}
	}
	for k, _ := range m {
		if ret > k {
			ret = k
		}
	}
	if ret == 1000000 {
		return -1
	}
	return ret
}