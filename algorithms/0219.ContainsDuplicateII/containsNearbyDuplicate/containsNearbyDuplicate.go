package containsNearbyDuplicate

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if m[v] != 0 && i-m[v]+1 <= k {
			return true
		}
		m[v] = i + 1
	}
	return false
}
