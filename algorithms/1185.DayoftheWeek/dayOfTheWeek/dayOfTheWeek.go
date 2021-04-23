package dayOfTheWeek

func dayOfTheWeek(day int, month int, year int) string {
	ret := dateToWeek(day, month, year)
	return ret
}

func dateToWeek(day int, month int, year int) string {
	var weekday = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	var y, m, c int
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}
	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1

	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	which_week := int(week)
	return weekday[which_week]
}
