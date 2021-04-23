package judgeCircle

func JudgeCircle(moves string) bool {
	x, y := 0, 0
	for _, v := range moves {
		if v == 'R' {
			x++
		} else if v == 'L' {
			x--
		} else if v == 'U' {
			y++
		} else {
			y--
		}
	}
	return x == 0 && y == 0
}
