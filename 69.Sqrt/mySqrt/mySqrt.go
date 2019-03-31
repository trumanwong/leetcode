package mySqrt

func MySqrt(x int) int {
	high := (x / 2) + 1
	low := 0
	for ; high >= low; {
		mid := (high + low) / 2
		if mid * mid == x {
			return mid
		} else if mid * mid > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}