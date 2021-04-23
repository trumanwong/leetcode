package readBinaryWatch

import "strconv"

func ReadBinaryWatch(num int) []string {
	ret := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if countOne(i)+countOne(j) == num {
				minute := strconv.Itoa(j)
				if j < 10 {
					minute = "0" + minute
				}
				hour := strconv.Itoa(i)
				ret = append(ret, hour+":"+minute)
			}
		}
	}
	return ret
}

func countOne(n int) int {
	ret := 0
	for n != 0 {
		n = n & (n - 1)
		ret++
	}
	return ret
}
