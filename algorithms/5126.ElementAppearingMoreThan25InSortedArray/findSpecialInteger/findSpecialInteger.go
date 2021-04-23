package findSpecialInteger

func findSpecialInteger(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
		if m[v] > int(float64(len(arr))*0.25) {
			return v
		}
	}
	return -1
}
