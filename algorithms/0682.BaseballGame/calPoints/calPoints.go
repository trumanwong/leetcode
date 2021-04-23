package calPoints

import "strconv"

func CalPoints(ops []string) int {
	res := 0
	nums := make([]int, 0)
	for _, op := range ops {
		if op == "C" && len(nums) > 0 {
			res -= nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		} else if op == "D" && len(nums) > 0 {
			res += nums[len(nums)-1] * 2
			nums = append(nums, nums[len(nums)-1]*2)
		} else if op == "+" {
			temp := 0
			for i := len(nums) - 1; i >= len(nums)-2 && i >= 0; i-- {
				temp += nums[i]
			}
			res += temp
			nums = append(nums, temp)
		} else {
			temp, _ := strconv.Atoi(op)
			res += temp
			nums = append(nums, temp)
		}
	}
	return res
}
