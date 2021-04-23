package licenseKeyFormatting

func LicenseKeyFormatting(S string, K int) string {
	res, temp := "", []byte{}
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] != '-' {
			val := S[i]
			if val >= 97 && val <= 122 {
				val -= 32
			}
			temp = append(temp, val)
		}
		if len(temp) == K {
			if res != "" {
				res += "-" + string(temp)
			} else {
				res += string(temp)
			}
			temp = []byte{}
		}
	}
	if len(temp) < K && len(temp) > 0 {
		if res != "" {
			res += "-" + string(temp)
		} else {
			res += string(temp)
		}
	}
	ans := []byte(res)
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}
