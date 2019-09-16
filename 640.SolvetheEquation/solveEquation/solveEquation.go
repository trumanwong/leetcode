package solveEquation

import "strconv"

func SolveEquation(equation string) string {
	bytes := []byte(equation)
	isLeft, isPlus := true, true
	start, x, val := -1, 0, 0
	for i, v := range bytes {
		if v == '=' || v == '+' || v == '-' {
			if start != -1 {
				temp, _ := strconv.Atoi(string(bytes[start:i]))
				if isLeft {
					if isPlus {
						val += temp
					} else {
						val -= temp
					}
				} else {
					if isPlus {
						val -= temp
					} else {
						val += temp
					}
				}
			}
			if v == '=' {
				isLeft, isPlus, start = false, true, -1
			} else if v == '+' {
				isPlus, start = true, -1
			} else {
				isPlus, start = false, -1
			}
		} else if v == 'x' {
			temp := 1
			if start != -1 {
				temp, _ = strconv.Atoi(string(bytes[start:i]))
			}
			if isLeft {
				if isPlus {
					x += temp
				} else {
					x -= temp
				}
			} else {
				if isPlus {
					x -= temp
				} else {
					x += temp
				}
			}
		} else if start == -1 {
			start = i
		}
	}
	if start != -1 {
		temp, _ := strconv.Atoi(string(bytes[start:len(bytes)]))
		if isPlus {
			val -= temp
		} else {
			val += temp
		}
	}
	if x == 0 && val == 0 {
		return "Infinite solutions"
	} else if x== 0 {
		return "No solution"
	} else {
		ret := strconv.Itoa(-val / x)
		return "x=" + ret
	}
}