package containsDuplicate

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = v
		}
	}
	return false
}
