package strangePrinter

import "math"

func strangePrinter(s string) int {
	f := make([][]int64, len(s))
	for i := range f {
		f[i] = make([]int64, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				f[i][j] = f[i][j-1]
			} else {
				f[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					f[i][j] = int64(math.Min(float64(f[i][j]), float64(f[i][k]+f[k+1][j])))
				}
			}
		}
	}
	return int(f[0][len(s)-1])
}
