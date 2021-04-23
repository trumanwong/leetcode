package queensAttacktheKing

func QueensAttacktheKing(queens [][]int, king []int) [][]int {
	res := make([][]int, 0)

	for _, queen := range queens {
		isR := true
		if queen[0] == king[0] {
			for _, v := range queens {
				if v[0] == king[0] && (v[1]-king[1])*(v[1]-queen[1]) < 0 {
					isR = false
					break
				}
			}

			if isR {
				res = append(res, queen)
			}
		} else if queen[1] == king[1] {
			for _, v := range queens {
				if v[1] == king[1] && (v[0]-king[0])*(v[0]-queen[0]) < 0 {
					isR = false
					break
				}
			}

			if isR {
				res = append(res, queen)
			}
		} else if queen[0]+queen[1] == king[0]+king[1] {
			for _, v := range queens {
				if v[0]+v[1] == king[0]+king[1] && (v[1]-king[1])*(v[1]-queen[1]) < 0 {
					isR = false
					break
				}
			}

			if isR {
				res = append(res, queen)
			}
		} else if queen[0]-queen[1] == king[0]-king[1] {
			for _, v := range queens {
				if v[0]-v[1] == king[0]-king[1] && (v[1]-king[1])*(v[1]-queen[1]) < 0 {
					isR = false
					break
				}
			}
			if isR {
				res = append(res, queen)
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
