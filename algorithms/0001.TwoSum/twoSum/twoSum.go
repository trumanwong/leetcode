package twoSum

func TwoSum(nums []int, target int) []int {
	m, res := make(map[int]int), make([]int, 0)
	for i, v := range nums {
		diff := target - v
		if _, ok := m[diff]; ok {
			res = append(res, []int{m[diff], i}...)
			break
		}
		m[v] = i
	}
	return res
}
