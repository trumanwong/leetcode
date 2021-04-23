package intersect

func Intersect(nums1 []int, nums2 []int) []int {
	m, res := make(map[int]int), []int{}
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
		}
		m[v]--
	}
	return res
}
