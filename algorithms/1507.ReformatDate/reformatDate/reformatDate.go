package reformatDate

import (
	"fmt"
	"strconv"
	"strings"
)

func ReformatDate(date string) string {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	s2month := make(map[string]int)
	for i := 0; i < 12; i++ {
		s2month[months[i]] = i + 1
	}
	arr := strings.Split(date, " ")
	day, _ := strconv.Atoi(arr[0][0 : len(arr[0])-2])
	return fmt.Sprintf("%s-%02d-%02d", arr[2], s2month[arr[1]], day)
}
