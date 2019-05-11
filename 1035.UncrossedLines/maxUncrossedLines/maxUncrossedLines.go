package maxUncrossedLines

func MaxUncrossedLines(A []int, B []int) int {
	n, m := len(A), len(B)
	MAXN := 555;
	a, b := make([]int, MAXN), make([]int, MAXN)
	var F [][]int
	F = make([][]int, MAXN)
	for i := 0; i < MAXN; i++ {
		F[i]= make([]int, MAXN)
	}

	for i := 0; i < n; i++ {
		a[i+1] = A[i]
	}
	for i := 0; i < m; i++ {
		b[i+1] = B[i]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i] == b[j] {
				F[i][j] = F[i-1][j-1] + 1
			} else {
				F[i][j] = max(F[i-1][j], F[i][j-1])
			}
		}
	}
	return F[n][m];
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}