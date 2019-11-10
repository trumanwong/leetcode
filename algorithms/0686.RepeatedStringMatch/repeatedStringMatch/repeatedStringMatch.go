package repeatedStringMatch

func RepeatedStringMatch(A string, B string) int {
	// 在A里去找B的首字母,找到就循环比较,最后返回循环的次数
	for i := 0; i < len(A); i++ {
		if A[i] != B[0] {
			continue
		}
		k, j := i, 0
		count := 1 // 循环次数
		for A[k] == B[j] {
			k++
			j++
			if j >= len(B) {
				return count
			}
			if k >= len(A) {
				k = 0
				count++
			}
		}
	}
	return -1
}