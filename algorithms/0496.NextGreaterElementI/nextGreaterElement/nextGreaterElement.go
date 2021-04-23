package nextGreaterElement

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	m, res := make(map[int]int), []int{}
	for i, v := range nums2 {
		m[v] = i
	}
	for _, v := range nums1 {
		temp := false
		if _, ok := m[v]; !ok {
			res = append(res, -1)
			continue
		}
		for j := m[v] + 1; j < len(nums2); j++ {
			if nums2[j] > v {
				res = append(res, nums2[j])
				temp = true
				break
			}
		}
		if !temp {
			res = append(res, -1)
		}
	}
	return res
}
