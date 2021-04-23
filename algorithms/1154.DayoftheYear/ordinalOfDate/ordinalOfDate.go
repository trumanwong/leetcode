package ordinalOfDate

import "strconv"

func OrdinalOfDate(date string) int {
	arr := []int{}
	start := -1
	s := []byte(date)
	for i, v := range s {
		if v == '-' {
			temp, _ := strconv.Atoi(string(s[start:i]))
			arr = append(arr, temp)
			start = -1
		} else if start == -1 {
			start = i
		}
	}

	if start != -1 {
		temp, _ := strconv.Atoi(string(s[start:len(s)]))
		arr = append(arr, temp)
	}
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (arr[0]%4 == 0 && arr[0]%100 != 0) || arr[0]%400 == 0 {
		months[1] = 29
	}

	ret := 0
	for i := 0; i < arr[1]-1; i++ {
		ret += months[i]
	}
	return ret + arr[2]
}
