package tribonacci

func Tribonacci(n int) int {
	t := []int{}
	t = append(t, []int{0, 1, 1}...)
	if n < 3 {
		return t[n]
	}
	t = append(t, t[0]+t[1]+t[2])
	for i := 4; i <= n; i++ {
		t = append(t, 0)
		for j := i - 3; j < i; j++ {
			t[i] += t[j]
		}
	}
	return t[n]
}
