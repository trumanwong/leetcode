package fib

func Fib(N int) int {
	if N < 2 {
		return N
	}
	ret := make([]int, N)
	ret[0] = 0
	ret[1] = 1
	for i := 2; i < N; i++ {
		ret[i] = ret[i-1] + ret[i-2]
	}
	return ret[N-1] + ret[N-2]
}
