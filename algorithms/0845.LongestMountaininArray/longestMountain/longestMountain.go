package longestMountain

func LongestMountain(A []int) int {
	start, end := 1, len(A)-2
	res := 0
	for i := start; i <= end; i++ {
		left, right := 0, 0
		prev := i
		for j := i - 1; j >= 0; j-- {
			if A[prev] > A[j] {
				left, prev = left+1, j
			} else {
				break
			}
		}
		prev = i
		for j := i + 1; j < len(A); j++ {
			if A[prev] > A[j] {
				right, prev = right+1, j
			} else {
				break
			}
		}
		if left > 0 && right > 0 {
			length := left + right + 1
			if length > res {
				res = length
			}
		}
	}
	return res
}
