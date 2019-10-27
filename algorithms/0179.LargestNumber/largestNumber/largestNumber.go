package largestNumber

import "strconv"

func LargestNumber(nums []int) string {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			s1, s2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
			if s1 + s2 > s2 + s1 {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	res := ""
	for i := len(nums) - 1; i >= 0; i-- {
		if len(res) == 0 && nums[i] == 0 {
			continue
		}
		res += strconv.Itoa(nums[i])
	}
	if len(res) == 0 {
		res = "0"
	}
	return res
}