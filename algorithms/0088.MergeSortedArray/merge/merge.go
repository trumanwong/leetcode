package merge

func Merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	m--
	n--
	for ; m >= 0 && n >= 0; p-- {
		if nums1[m] > nums2[n] {
			nums1[p] = nums1[m]
			m--
		} else {
			nums1[p] = nums2[n]
			n--
		}
	}
	for ; n >= 0; n-- {
		nums1[p] = nums2[n]
		p--
	}
}
