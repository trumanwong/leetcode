package combine

func Combine(n int, k int) [][]int {
	nums := make([]int, k+1)
	for i := 0; i < k; i++ {
		nums[i] = i + 1
	}
	nums[k] = n + 1

	res, j := make([][]int, 0), 0
	for j < k {
		temp := make([]int, k)
		copy(temp, nums[:k])
		res = append(res, temp)
		j = 0
		for j < k && nums[j+1] == nums[j]+1 {
			nums[j] = j + 1
			j++
		}
		nums[j] += 1
	}
	return res
}
