package invalidTransactions

import (
	"strconv"
	"strings"
)

func InvalidTransactions(transactions []string) []string {
	if len(transactions) == 0 {
		return []string{}
	}

	ts := make([][]string, 0)
	for _, transaction := range transactions {
		temp := strings.Split(transaction, ",")
		temp = append(temp, transaction)
		ts = append(ts, temp)
	}

	m := make(map[string]int)
	for _, v := range ts {
		amount, _ := strconv.Atoi(v[2])
		if amount > 1000 {
			m[v[4]]++
		}
		time1, _ := strconv.Atoi(v[1])
		for _, o := range ts {
			time2, _ := strconv.Atoi(o[1])
			if v[0] == o[0] && abs(time1-time2) <= 60 && v[3] != o[3] {
				m[v[4]]++
			}
		}
	}

	ret := make([]string, 0)
	for k, _ := range m {
		ret = append(ret, k)
	}

	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
