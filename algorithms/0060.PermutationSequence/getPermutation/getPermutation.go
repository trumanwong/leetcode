package getPermutation

func GetPermutation(n int, k int) string {
	data := []byte("0123456789")

	f := make([]int, n + 1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = f[i - 1] * i
	}

	res := ""
	for k >= 1 && n != 0 {
		pos := k / f[n - 1]
		if k % f[n - 1] != 0 {
			pos++
		}

		k = k - (pos - 1) * f[n - 1]
		res = res + string(data[pos])
		data = append(data[0:pos], data[pos + 1:]...)
		n--
	}
	return res
}