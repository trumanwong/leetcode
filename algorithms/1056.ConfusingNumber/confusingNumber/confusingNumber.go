package confusingNumber

import (
	"strconv"
)

func ConfusingNumber(N int) bool {
	m := make(map[byte]byte)
	m['0'] = '0'
	m['1'] = '1'
	m['6'] = '9'
	m['8'] = '8'
	m['9'] = '6'
	n := []byte(strconv.Itoa(N))
	temp := make([]byte, len(n))
	for i, v := range n {
		if _, ok := m[v]; !ok {
			return false
		}
		temp[len(n)-1-i] = m[n[i]]
	}
	if string(temp) != string(n) {
		return true
	}
	return false
}
