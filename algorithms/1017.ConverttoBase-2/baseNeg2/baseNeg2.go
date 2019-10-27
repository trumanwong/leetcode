package baseNeg2

func BaseNeg2(N int) string {
	res, flag := "", true
	if N == 0 {
		return "0"
	}

	for N != 0 {
		if N & 1 != 0 {
			if flag {
				N -= 1
			} else {
				N += 1
			}
			res = "1" + res
		} else {
			res = "0" + res
		}
		N /= 2
		flag = !flag
	}
	return res
}