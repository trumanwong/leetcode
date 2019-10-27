package maxArea

func MaxArea(height []int) int {
	if len(height) <= 1 {
		return -1
	}
	left, right := 0, len(height) - 1
	res := 0
	for left < right {
		h := min(height[left], height[right])
		res = max(res, h * (right - left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}