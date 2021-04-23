package twoSum

func TwoSum(numbers []int, target int) []int {
	start, end, res := 0, len(numbers)-1, []int{}
	for start < end {
		temp := numbers[start] + numbers[end]
		if temp == target {
			res = append(res, []int{start + 1, end + 1}...)
			break
		} else if temp > target {
			end--
		} else {
			start++
		}
	}
	return res
}
