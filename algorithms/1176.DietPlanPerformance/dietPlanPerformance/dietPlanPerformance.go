package dietPlanPerformance

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	count, sum, ret := 0, 0, 0
	for i, v := range calories {
		count++
		sum += v
		if count == k {
			if sum < lower {
				ret -= 1
			} else if sum > upper {
				ret += 1
			}
			sum -= calories[i-k+1]
			count--

			if i == len(calories)-1 {
				count = 0
			}
		}
	}
	if count > 0 {
		if sum > upper {
			ret += 1
		} else if sum < lower {
			ret -= 1
		}
	}
	return ret
}
