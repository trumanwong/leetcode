package selfDividingNumbers

func SelfDividingNumbers(left int, right int) []int {
	ret := []int{}
	for i := left; i <= right; i++ {
		if i % 10 == 0 {
			continue
		}
		if i < 10 {
			ret = append(ret, i)
		} else {
			temp := i
			b := false
			for ; temp % 10 > 0; {
				if temp % 10 == 0 || i % (temp % 10) != 0 {
					b = false
					break
				} else {
					b = true
				}
				temp = temp / 10
				if temp >= 10 && temp % 10 == 0 {
					b = false
					break
				}
			}
			if b {
				ret = append(ret, i)
			}
		}
	}
	return ret
}