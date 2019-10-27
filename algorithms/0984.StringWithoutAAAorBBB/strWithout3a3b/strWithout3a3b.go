package strWithout3a3b

func StrWithout3a3b(A int, B int) string {
	ans := ""

	for A > 0 || B > 0 {
		writeA := false
		L := len(ans)
		if L >= 2 && ans[L - 1] == ans [L - 2] {
			if string(ans[L - 1]) == "b" {
				writeA = true
			}
		} else if A >= B {
			writeA = true
		}

		if writeA {
			A--
			ans += "a"
		} else {
			B--
			ans += "b"
		}
	}
	return ans
}