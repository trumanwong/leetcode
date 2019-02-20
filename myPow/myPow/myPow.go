package myPow

/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/
func fast_pow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	//计算一半的pow值
	p := MyPow(x, n / 2)
	//如果n是奇数，必然少算了一个x，需要乘以x
	if n % 2 == 1 {
		return p * p * x
	} else {
		return p * p
	}
}

func MyPow(x float64, n int) float64 {
	//正负处理
	if n < 0 {
		return 1 / fast_pow(x, -n)
	} else {
		return fast_pow(x, n)
	}
}
