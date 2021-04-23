package earliestAcq

func EarliestAcq(logs [][]int, N int) int {
	for i := 0; i < len(logs); i++ {
		for j := i + 1; j < len(logs); j++ {
			if logs[j][0] < logs[i][0] {
				logs[i], logs[j] = logs[j], logs[i]
			}
		}
	}

	arr := [][]int{}

	for i := 0; i < len(logs); i++ {
		v1, v2 := logs[i][1], logs[i][2]
		if len(arr) == 0 {
			arr = append(arr, []int{v1, v2})
		} else {
			has := false
			for j := 0; j < len(arr); j++ {
				if inArray(v1, arr[j]) || inArray(v2, arr[j]) {
					has = true
					if inArray(v1, arr[j]) && !inArray(v2, arr[j]) {
						arr[j] = append(arr[j], v2)
					} else if inArray(v2, arr[j]) && !inArray(v1, arr[j]) {
						arr[j] = append(arr[j], v1)
					}
					break
				}
			}
			if !has {
				arr = append(arr, []int{v1, v2})
			}
		}

		for j := 0; j < len(arr); j++ {
			has, n := false, -1
			for k := j + 1; k < len(arr); k++ {
				if array_intersect(arr[k], arr[j]) {
					for m := 0; m < len(arr[k]); m++ {
						if !inArray(arr[k][m], arr[j]) {
							arr[j] = append(arr[j], arr[k][m])
						}
					}
					n = k
					has = true
					break
				}
			}
			if has {
				arr = append(arr[:n], arr[n+1:]...)
				break
			}
		}
		if len(arr[0]) >= N {
			return logs[i][0]
		}
	}
	return -1
}

func array_intersect(a []int, b []int) bool {
	for _, v := range a {
		if inArray(v, b) {
			return true
		}
	}
	return false
}

func inArray(needle int, stack []int) bool {
	for _, v := range stack {
		if v == needle {
			return true
		}
	}
	return false
}
