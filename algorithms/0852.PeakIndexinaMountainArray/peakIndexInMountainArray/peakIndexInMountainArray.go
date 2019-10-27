package peakIndexInMountainArray

func PeakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A); i++ {
		if A[i] < A[i - 1] {
			return i - 1
		}
	}
	return -1
}