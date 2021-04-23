package majorityElement

func MajorityElement(nums []int) []int {
	m, judge, res := make(map[int]int), len(nums)/3, []int{}
	for _, v := range nums {
		m[v]++
	}
	for i, v := range m {
		if v > judge {
			res = append(res, i)
		}
	}
	return res
}
