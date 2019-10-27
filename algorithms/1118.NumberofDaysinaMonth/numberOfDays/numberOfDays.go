package numberOfDays

func NumberOfDays(Y int, M int) int {
	ret := 0
	if M == 2 {
		if isLeapYear(Y) {
			ret = 29
		} else {
			ret = 28
		}
	} else if M == 4 || M == 6 || M == 9 || M == 11 {
		ret = 30
	} else {
		ret = 31
	}
	return ret
}

func isLeapYear(year int) bool {
	//判断是否为闰年
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return true
	}
	return false
}