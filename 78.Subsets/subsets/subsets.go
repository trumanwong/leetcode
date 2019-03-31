package subsets

var res [][]int
func Subsets(nums []int) [][]int {
	res = [][]int{}
	dfs(nums, []int{}, 0)
	return res
}

func dfs(nums []int, temp []int, index int) {
	t := make([]int, len(temp))
	copy(t, temp)
	res = append(res, t)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		dfs(nums, temp, i + 1)
		temp = temp[:len(temp) - 1]
	}
}