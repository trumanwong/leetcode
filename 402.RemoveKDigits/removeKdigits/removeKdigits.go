package removeKdigits

func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}

	ret := []byte(num)
	for i := 0; i < k; i++ {
		idx := 0
		for j := 1; j < len(ret) && ret[j] >= ret[j - 1]; j++ {
			idx = j
		}
		ret = append(ret[0:idx], ret[idx + 1:]...)
		for len(ret) > 1 && ret[0] == '0' {
			ret = ret[1:]
		}
	}
	return string(ret)
}