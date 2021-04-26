package shipWithinDays

import "sort"

func ShipWithinDays(weights []int, D int) int {
	// 确定二分查找左右边界
	l, r := 0, 0
	for _, v := range weights {
		if v > l {
			l = v
		}
		r += v
	}

	return l + sort.Search(r-l, func(x int) bool {
		x += l
		day := 1 // 需要运送的天数
		sum := 0 // 当前这一天已经运送的包裹重量之和
		for _, v := range weights {
			if sum+v > x {
				day++
				sum = 0
			}
			sum += v
		}
		return day <= D
	})
}
