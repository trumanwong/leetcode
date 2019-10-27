package missingNumber

func MissingNumber(arr []int) int {
	diff := (arr[len(arr) - 1] - arr[0]) / len(arr)
	if diff == 0 {
		return arr[0]
	}
	next := arr[0] + diff
	for i := 1; i < len(arr); i++ {
		if next != arr[i] {
			return next
		}
		next += diff
	}
	return 0
}