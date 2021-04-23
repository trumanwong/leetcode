package climbStairs

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev, after, res := 1, 2, 0
	for i := 2; i < n; i++ {
		res = prev + after
		prev, after = after, res
	}
	return res
}
