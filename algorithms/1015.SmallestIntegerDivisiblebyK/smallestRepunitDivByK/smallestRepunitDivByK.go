package smallestRepunitDivByK

import "strconv"

func SmallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	kStr := strconv.Itoa(K)
	res := ""
	for i := 0; i < len(kStr); i++ {
		res += "1"
	}
	b, _ := strconv.Atoi(res)
	if b < K {
		res += "1"
		b, _ = strconv.Atoi(res)
	}
	for b%K != 0 {
		res += "1"
		b %= K
		temp := strconv.Itoa(b) + "1"
		b, _ = strconv.Atoi(temp)
	}
	return len(res)
}
