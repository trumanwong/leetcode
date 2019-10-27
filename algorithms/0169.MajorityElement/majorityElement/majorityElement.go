package majorityElement

func MajorityElement(nums []int) int {
	m := make(map[int]int)
	max, res := 0, 0
	for _, v := range nums {
		m[v]++
		if m[v] > max {
			max = m[v]
			res = v
		}
	}
	return res
}