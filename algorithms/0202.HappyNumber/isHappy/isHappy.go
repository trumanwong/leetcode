package isHappy

func IsHappy(n int) bool {
	m := make(map[int]int)
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n]++
		temp, sum := n, 0
		for temp > 0 {
			sum += (temp % 10) * (temp % 10)
			temp /= 10
		}
		n = sum
	}
	return true
}
