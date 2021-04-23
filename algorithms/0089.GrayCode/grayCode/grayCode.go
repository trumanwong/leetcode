package grayCode

func GrayCode(n int) []int {
	length := 1 << uint(n)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = (i ^ i>>1)
	}
	return res
}
