package prefixesDivBy5

func PrefixesDivBy5(A []int) []bool {
	res, sum := make([]bool, len(A)), 0
	for i, v := range A {
		sum = (sum + v) * 2 % 10
		if sum%5 == 0 {
			res[i] = true
		}
	}
	return res
}
