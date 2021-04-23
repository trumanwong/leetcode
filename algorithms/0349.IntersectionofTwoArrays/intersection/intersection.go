package intersection

func Intersection(nums1 []int, nums2 []int) []int {
	m, res := make(map[int]int), []int{}
	for _, v := range nums1 {
		m[v] = 1
	}

	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}
