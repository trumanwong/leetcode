package findMedianSortedArrays

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	l, r, mid := 0, m, (m + n + 1) / 2
	for l <= r {
		i := (l + r) / 2
		j := mid - i
		if i < r && nums2[j - 1] > nums1[i] {
			l = i + 1 // i is too small
		} else if i > l && nums1[i - 1] > nums2[j] {
			r = i - 1 // i is too big
		} else {
			var maxLeft float64
			if i == 0 {
				maxLeft = float64(nums2[j - 1])
			} else if j == 0 {
				maxLeft = float64(nums1[i - 1])
			} else {
				maxLeft = math.Max(float64(nums1[i - 1]), float64(nums2[j - 1]))
			}

			if (m + n) % 2 == 1 {
				return maxLeft
			}

			var minRight float64
			if i == m {
				minRight = float64(nums2[j])
			} else if  j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums2[j]), float64(nums1[i]))
			}
			return (maxLeft + minRight) / 2
		}
	}
	return 0
}