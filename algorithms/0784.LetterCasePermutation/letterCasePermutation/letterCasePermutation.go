package letterCasePermutation

func LetterCasePermutation(S string) []string {
	res := []string{}
	var DFS func(i int, sByte []byte)
	DFS = func(i int, sByte []byte) {
		if i >= len(sByte) {
			res = append(res, string(sByte))
			return
		}

		if sByte[i] >= 'A' && sByte[i] <= 'Z' {
			sByte[i] += 32
			DFS(i + 1, sByte)
			sByte[i] -= 32
			DFS(i + 1, sByte)
		} else if sByte[i] >= 'a' && sByte[i] <= 'z' {
			sByte[i] -= 32
			DFS(i + 1, sByte)
			sByte[i] += 32
			DFS(i + 1, sByte)
		} else {
			// 数字跳过
			DFS(i + 1, sByte)
		}
	}
	DFS(0, []byte(S))
	return res
}