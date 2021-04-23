package sampleStats

func SampleStats(count []int) []float64 {
	min, max, sum, mid, most := 256, -1, 0, float64(0), 0
	most_v, indexs := -1, 0
	arr, m := []int{}, make(map[int]int)
	for i, v := range count {
		if v != 0 {
			arr, m[i] = append(arr, i), v
			indexs += v
			sum += i * v
			if i < min {
				min = i
			}

			if i > max {
				max = i
			}

			if most_v < v {
				most, most_v = i, v
			}
		}
	}
	avg := float64(sum) / float64(indexs)
	unique, left, right, is_mid := 0, 0, 0, false
	if indexs%2 != 0 {
		unique = indexs / 2
		is_mid = true
	} else {
		left, right = indexs/2, indexs/2+1
	}
	indexs, sum_mid := 0, 0
	for _, i := range arr {
		indexs += m[i]
		if is_mid {
			if unique <= indexs {
				mid = float64(i)
				break
			}
		} else {
			if left != -1 && left <= indexs {
				sum_mid, left = sum_mid+i, -1
			}
			if right != -1 && right <= indexs {
				sum_mid, right = sum_mid+i, -1
			}
			if left == -1 && right == -1 {
				break
			}
		}
	}
	if left == -1 && right == -1 {
		mid = float64(sum_mid) / float64(2)
	}
	return []float64{float64(min), float64(max), avg, mid, float64(most)}
}
