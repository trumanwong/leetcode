package pivotIndex

func PivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, v := range nums {
		sum += v
	}

	for i, v := range nums {
		rightSum := sum - v - leftSum
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}
	return -1
}