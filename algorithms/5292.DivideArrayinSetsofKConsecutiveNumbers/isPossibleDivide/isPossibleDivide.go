package isPossibleDivide

func IsPossibleDivide(nums []int, k int) bool {
	if len(nums) % k != 0 {
		return false
	}

	m, arr := make(map[int]int), make([]int, 0)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			arr = append(arr, v)
		}
		m[v]++
	}
	index, last := 0, 0
	for len(arr) != 0 {
		if index >= k {
			index, last = 0, 0
		}
		if last >= arr[index] {
			return false
		}
		last = arr[index]

		m[arr[index]]--
		if m[arr[index]] == 0 {
			delete(m, arr[index])
			arr = append(arr[:index], arr[index + 1:]...)
		}
		index++
		if k > len(arr) {
			index = 0
		}
	}
	return true
}