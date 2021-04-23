package countSteppingNumbers

import (
	"sort"
	"strconv"
)

func CountSteppingNumbers(low int, high int) []int {
	res := make([]int, 0)
	ls, hs := strconv.Itoa(low), strconv.Itoa(high)
	var helper func(curr int, n int)
	helper = func(curr int, n int) {
		if n == 0 {
			if curr >= low && curr <= high {
				res = append(res, curr)
			}

			return
		}
		last := curr % 10
		if last+1 <= 9 {
			helper(curr*10+last+1, n-1)
		}
		if last-1 >= 0 {
			helper(curr*10+last-1, n-1)
		}
	}
	for i := len(ls); i <= len(hs); i++ {
		start := 0
		if i > 1 {
			start = 1
		}
		for j := start; j <= 9; j++ {
			helper(j, i-1)
		}
	}
	sort.Ints(res)
	return res
}
