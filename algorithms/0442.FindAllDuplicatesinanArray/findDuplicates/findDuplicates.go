package findDuplicates

func FindDuplicates(nums []int) []int {
	ret := []int{}
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			ret = append(ret, num)
		} else {
			m[num] = num
		}
	}
	return ret
}
