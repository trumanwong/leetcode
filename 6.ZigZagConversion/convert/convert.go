package convert

func Convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || length <= numRows {
		return s
	}

	res := ""
	for i := 0; i < numRows; i++ {
		j := i
		res += string(s[j])
		// 两段跳
		for ; j < length; {
			p := 2 * (numRows - i -1)
			q := 2 * i

			j += p
			// 红色第一跳
			if p > 0 && j < length {
				res += string(s[j])
			}
			j += q
			// 蓝色第二跳
			if q > 0 && j < length {
				res += string(s[j])
			}
		}
	}
	return res
}