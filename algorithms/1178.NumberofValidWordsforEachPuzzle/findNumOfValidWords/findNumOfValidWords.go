package findNumOfValidWords

func FindNumOfValidWords(words []string, puzzles []string) []int {
	bit2cnt := make(map[int]int) // "字符串word转化为一个数(位运算): 该模式出现的次数"  的映射
	// 统计每个word模式的数量, map形式保存, 位运算进行操作
	for i := 0; i < len(words); i++ {
		k := 0
		for j := 0; j < len(words[i]); j++ {
			k |= (1 << (words[i][j] - 'a'))
		}
		bit2cnt[k] += 1
	}

	res := make([]int, 0)
	for i := 0; i < len(puzzles); i++ {
		res = append(res, 0)
		k := 0
		for j := 0; j < len(puzzles[i]); j++ {
			k |= (1 << (puzzles[i][j] - 'a'))
		}

		for j := k; j != 0; j = (j - 1) & k {
			if (1 << (puzzles[i][0] - 'a') & j) != 0 {
				res[i] += bit2cnt[j]
			}
		}
	}
	return res
}
