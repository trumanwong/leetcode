package carPooling

func CarPooling(trips [][]int, capacity int) bool {
	if len(trips) == 0 {
		return true
	}
	for i := 0; i < len(trips); i++ {
		for j := i + 1; j < len(trips); j++ {
			if trips[j][1] < trips[i][1] {
				trips[i], trips[j] = trips[j], trips[i]
			}
		}
	}
	remain := capacity - trips[0][0]
	out := make(map[int]int)
	out[trips[0][2]] = trips[0][0]
	for i := 1; i < len(trips); i++ {
		remain -= trips[i][0]
		for k, v := range out {
			if k <= trips[i][1] {
				remain += v
				delete(out, k)
			}
		}
		if remain < 0 {
			return false
		}
		out[trips[i][2]] += trips[i][0]
	}
	if remain < 0 {
		return false
	}
	return true
}