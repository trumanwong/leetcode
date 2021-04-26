package mySqrt

func MySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		mid := (l + r) / 2
		if x / mid == mid {
			return mid
		} else if x / mid < mid {
			r = mid - 1
		} else {
			l =  mid + 1
		}
	}
	return r
}