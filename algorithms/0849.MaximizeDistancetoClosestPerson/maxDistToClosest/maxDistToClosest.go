package maxDistToClosest

func MaxDistToClosest(seats []int) int {
	/*
	   第一个 1 到最左边的距离
	   相邻的 1 中间位置到彼此的距离
	   最后一个 1 到最右边的距离
	*/
	res, last := 1, -1
	for i, v := range seats {
		if v != 1 {
			continue
		}

		if last == -1 {
			res = max(res, i)
		} else {
			res = max(res, (i-last)/2)
		}
		last = i
	}
	return max(res, len(seats)-1-last)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
