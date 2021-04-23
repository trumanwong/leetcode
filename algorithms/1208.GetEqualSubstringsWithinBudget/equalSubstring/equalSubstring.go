package equalSubstring

func EqualSubstring(s string, t string, maxCost int) int {
	diff := make([]int, len(s))
	if len(s) > len(t) {
		diff = make([]int, len(t))
	}

	i := 0
	ret, count := 0, 0
	for i < len(s) && i < len(t) {
		diff[i] = abs(int(s[i]) - int(t[i]))
		i++
	}

	temp := make([]int, len(diff))
	temp[0], count, ret = diff[0], 1, 1
	if temp[0] > maxCost {
		temp[0], count, ret = 0, 0, 0
	}
	for i := 1; i < len(diff); i++ {
		if temp[i-1]+diff[i] <= maxCost {
			temp[i] = temp[i-1] + diff[i]
			count++
		} else if diff[i] <= maxCost {
			temp[i] = temp[i-1] + diff[i]
			for temp[i] > maxCost {
				temp[i] -= diff[i-count]
				count--
			}
			count++
		} else {
			temp[i], count = 0, 0
		}

		if count > ret {
			ret = count
		}
	}

	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
