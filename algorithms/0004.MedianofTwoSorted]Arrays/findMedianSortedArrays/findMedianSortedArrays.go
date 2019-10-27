package findMedianSortedArrays

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	A, B := make([]int, len(nums1)), make([]int, len(nums2))
	copy(A, nums1)
	copy(B, nums2)
	if m > n {
		A, B = make([]int, len(nums2)), make([]int, len(nums1))
		copy(A, nums2)
		copy(B, nums1)
		m, n = n, m
	}

	l, r, mid := 0, m, (m + n + 1) / 2
	for l <= r {
		i := (l + r) / 2
		j := mid - i
		if i < r && B[j - 1] > A[i] {
			l = i + 1 // i is too small
		} else if i > l && A[i - 1] > B[j] {
			r = i - 1 // i is too big
		} else {
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(B[j - 1])
			} else if j == 0 {
				maxLeft = float64(A[i - 1])
			} else {
				maxLeft = math.Max(float64(A[i - 1]), float64(B[j - 1]))
			}

			if (m + n) / 2 == 1 {
				return maxLeft
			}

			var minRight float64
			if i == m {
				minRight = float64(B[j])
			} else if  j == n {
				minRight = float64(A[i])
			} else {
				minRight = math.Min(float64(B[j]), float64(A[i]))
			}
			return (maxLeft + minRight) / 2
		}
	}
	return 0
}