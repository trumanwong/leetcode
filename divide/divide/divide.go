package divide

import "math"

func Divide(dividend int, divisor int) int {
	if dividend == divisor {
		return 1
	}
	if dividend == -divisor || -dividend == divisor {
		return -1
	}
	if divisor == 1 {
		return dividend
	}

	sign := true
	ans := 0

	if dividend > 0 {
		if divisor < 0 {
			if dividend + divisor < 0 {
				return 0
			}
			sign = false
			divisor = -divisor
		}
	} else if divisor > 0 {
		if dividend + divisor > 0 {
			return 0
		}
		sign = false
		dividend = -dividend
	} else {
		if dividend - divisor > 0 {
			return 0
		}
		sign = true
		dividend = -dividend
		divisor = -divisor
	}

	var x, y int
	for ; dividend > divisor; {
		x = divisor
		y = 1
		for ; dividend >= x << 1; {
			x <<= 1
			y <<= 1
		}
		dividend -= x
		ans += y
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	} else if ans < math.MinInt32 {
		ans = math.MaxInt32
	}
	if sign {
		return ans
	} else {
		return -ans
	}
}
