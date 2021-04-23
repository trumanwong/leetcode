package checkRecord

func CheckRecord(s string) bool {
	acount, lcount := 0, 0
	for _, v := range s {
		if v == 'A' {
			acount++
		}
		if v == 'L' {
			lcount++
		} else {
			lcount = 0
		}
		if lcount > 2 {
			return false
		}
	}
	return acount <= 1
}
