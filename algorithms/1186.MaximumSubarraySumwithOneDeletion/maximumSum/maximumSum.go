package maximumSum

func maximumSum(arr []int) int {
	left, right := make([]int, len(arr)), make([]int, len(arr))
	left[0], right[len(arr)-1] = arr[0], arr[len(arr)-1]
	ret := max(left[0], right[len(arr)-1])
	for i := 1; i < len(arr); i++ {
		left[i] = arr[i] + max(0, left[i-1])
		ret = max(ret, left[i])
	}

	for i := len(arr) - 2; i >= 0; i-- {
		right[i] = arr[i] + max(0, right[i+1])
		ret = max(ret, right[i])
	}

	for i := 1; i <= len(arr)-2; i++ {
		ret = max(left[i-1]+right[i+1], ret)
	}

	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
