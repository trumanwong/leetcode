package nextGreaterElements

func NextGreaterElements(nums []int) []int {
	res := []int{}
	for i, v := range nums {
		temp := false
		j := i + 1
		if j == len(nums) {
			j = 0
		}
		for {
			if i == j {
				break
			}

			if nums[j] > v {
				res = append(res, nums[j])
				temp = true
				break
			}

			j++
			if j == len(nums) {
				j = 0
			}
		}
		if !temp {
			res = append(res, -1)
		}
	}
	return res
}