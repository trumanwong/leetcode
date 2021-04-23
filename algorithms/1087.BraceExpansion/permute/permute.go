package permute

func Permute(S string) []string {
	left, right, s := false, false, []byte(S)
	res, index := []string{}, 0
	m := make(map[int][]byte)
	for _, v := range s {
		if v == '{' {
			left = true
			right = false
		} else if v == '}' {
			right = true
			left = false
			index++
		} else if v == ',' {
			continue
		} else {
			m[index] = append(m[index], v)
			if !left || right {
				index++
			}
		}
	}

	var recursive func(bytes []byte, i int)
	recursive = func(bytes []byte, im int) {
		if im == index {
			res = append(res, string(bytes))
			return
		}
		for j := 0; j < len(m[im]); j++ {
			bytes = append(bytes, m[im][j])
			recursive(bytes, im+1)
			bytes = bytes[:len(bytes)-1]
		}
	}
	recursive([]byte{}, 0)
	for i := 0; i < len(res); i++ {
		for j := i; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return res
}
