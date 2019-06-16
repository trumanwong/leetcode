package rearrangeBarcodes

func RearrangeBarcodes(barcodes []int) []int {
	if len(barcodes) < 3 {
		return barcodes
	}
	
	m, length := make(map[int]int), len(barcodes)
	res := make([]int, length)
	for _, v := range barcodes {
		m[v]++
	}
	
	if len(barcodes) == 3 {
		for k, v := range m {
			if v == 2 {
				res[0], res[2] = k, k
			} else {
				res[1] = k
			}
		}
		return res
	}

	step := 0
	for k, v := range m {
		temp, i := v, 0
		for res[i + step] != 0 {
			i++
		}
		for temp > 0 {
			if i + step >= length {
				res[i] = k
			} else {
				res[i+step] = k
			}
			temp--
			i += 2
		}
		step++
	}
	return res
}