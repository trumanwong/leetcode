package brokenCalc

func BrokenCalc(X int, Y int) int {
	res := 0
	for Y > X {
		res++
		if Y%2 == 0 {
			Y /= 2
		} else {
			Y++
		}
	}
	return res + X - Y
}
