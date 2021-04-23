package findNthDigit

import (
	"math"
	"strconv"
)

func FindNthDigit(n int) int {
	/*
	   1、1-9有9个数，10-99有2*10*9个数字，100-999有3*100*9个数字，1000-9999有4*1000*9个数字，以此类推
	   2、设置一个标志位i，每一个区间都有固定的标志位，例如1-9是1，10-99是2，以此类推，然后用n减去每个取键得值，确定n在哪个区间
	   3、在得到区间中确定的数字，将其变为string，然后得到确定的数字
	*/
	i := 1
	for n > i*int(math.Pow10(i-1))*9 {
		n -= i * int(math.Pow10(i-1)*9) // 小于区间的值要减去，直到得到确定的区间
		i++                             // 标志位++
	}

	am_num := (n-1)/i + int(math.Pow10(i-1)) // 确定区间中数字
	num := strconv.Itoa(am_num)
	if n%i == 0 {
		res, _ := strconv.Atoi(string(num[i-1]))
		return res
	}
	res, _ := strconv.Atoi(string(num[n%i-1]))
	return res
}
