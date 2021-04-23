package smallestDivisor

func smallestDivisor(nums []int, threshold int) int {
	l, r := 0, 1000001
	res := r
	for l <= r {
		m := (l + r) / 2
		sum := 0
		for _, v := range nums {
			sum += v / m
			if v%m != 0 {
				sum += 1
			}
		}

		if sum <= threshold {
			res = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return res
}
