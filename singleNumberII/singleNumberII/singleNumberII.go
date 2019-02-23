package singleNumberII

func SingleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a = ^b & (a ^ num)
		b = ^a & (b ^ num)
	}
	return a
}