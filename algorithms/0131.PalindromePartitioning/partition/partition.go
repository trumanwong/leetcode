package partition

func Partition(s string) [][]string {
	result := make([][]string, 0)
	f := make([][]bool, len(s))
	for i := range f {
		f[i] = make([]bool, len(s))
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			f[i][j] = s[i] == s[j] && f[i + 1][j - 1]
		}
	}

	splits := make([]string, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i == len(s) {
			result = append(result, append([]string(nil), splits...))
			return
		}
		for j := i; j < len(s); j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return result
}