package canMeasureWater

func CanMeasureWater(x int, y int, z int) bool {
	if x + y < z {
		return false
	}

	v := gcd(x, y)
	if v == 0 || z % v == 0 {
		return true
	}
	return false
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x % y)
}