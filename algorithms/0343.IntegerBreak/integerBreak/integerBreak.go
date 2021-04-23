package integerBreak

func IntegerBreak(n int) int {
	// 求函数y=(n/x)^x(x>0)的最大值，可得x=e时y最大，但只能分解成整数，故x取2或3，由于6=2+2+2=3+3，显然2^3=8<9=3^2,故应分解为多个3
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	res := 1
	for n > 4 {
		n = n - 3
		res = res * 3
	}
	res = res * n
	return res
}
