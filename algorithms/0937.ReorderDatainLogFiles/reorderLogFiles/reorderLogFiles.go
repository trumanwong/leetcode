package reorderLogFiles

func ReorderLogFiles(logs []string) []string {
	nums, non_nums := []string{}, []string{}
	ks, vs := []string{}, []string{}

	for _, log := range logs {
		start := -1

		for i, v := range log {
			if v == ' ' && start == -1 {
				start = i
			} else if start != -1 && v != ' ' {
				if v >= '0' && v <= '9' {
					nums = append(nums, log)
				} else {
					k1, v1 := getKV([]byte(log))
					ks, vs = append(ks, k1), append(vs, v1)
					non_nums = append(non_nums, log)
				}
				break
			}
		}
	}

	for i := 0; i < len(non_nums); i++ {
		for j := i + 1; j < len(non_nums); j++ {
			judge1, judge2 := false, false

			if vs[i] == vs[j] {
				temp_i, temp_j := []byte(ks[i]), []byte(ks[j])

				for t := 0; t < len(ks[i]) && t < len(ks[j]); t++ {
					if temp_i[t] > temp_j[t] {
						judge1 = true
						break
					}
				}
			} else {
				temp_i, temp_j := []byte(vs[i]), []byte(vs[j])
				for t := 1; t < len(vs[i]) && t < len(vs[j]); t++ {
					if temp_i[t] > temp_j[t] {
						judge2 = true
						break
					} else if temp_i[t] < temp_j[t] {
						break
					}
				}
			}
			if judge1 || judge2 {
				ks[i], ks[j] = ks[j], ks[i]
				vs[i], vs[j] = vs[j], vs[i]
				non_nums[i], non_nums[j] = non_nums[j], non_nums[i]
			}
		}
	}

	ret := []string{}
	ret = append(ret, non_nums...)
	ret = append(ret, nums...)
	return ret
}

func getKV(value []byte) (string, string) {
	rk, rv := []byte{}, []byte{}
	start := -1
	for i, v := range value {
		if start != -1 && v != ' ' {
			rv = append(rv, value[start:len(value)]...)
			break
		} else if start == -1 && v == ' ' {
			start = i
			rk = append(rk, value[0:start]...)
		}
	}
	return string(rk), string(rv)
}
