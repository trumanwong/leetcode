package ReformatDate

import (
	"strconv"
	"strings"
)

func ReformatDate(date string) string {
	days := []string{"1st", "2nd", "3rd", "4th", "5th", "6th", "7th","8th","9th","10th","11th","12th","13th","14th","15th","16th","17th","18th","19th","20th","21th","22th","23th","24th","25th","26th","27th","28th","29th","30th", "31st"}
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	arr := strings.Split(date, " ")
	day, month := 0, 0
	for i := 0; i < len(days); i++ {
		if arr[0] == days[i] {
			day = i + 1
		}
	}
	for i := 0; i < len(months); i++ {
		if arr[1] == months[i] {
			month = i + 1
		}
	}
	res := arr[2]
	res += "-"
	if month < 10 {
		res += "0" + strconv.Itoa(month)
	} else {
		res += strconv.Itoa(month)
	}
	res += "-"
	if day < 10 {
		res += "0" + strconv.Itoa(day)
	} else {
		res += strconv.Itoa(day)
	}
	return res
}