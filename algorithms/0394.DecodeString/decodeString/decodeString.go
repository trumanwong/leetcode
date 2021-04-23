package decodeString

import "strconv"

func DecodeString(s string) string {
	res, cur := "", ""
	nums, strs := make([]int, 0), make([]string, 0)
	num := 0
	for _, v := range s {
		if v >= '0' && v <= '9' {
			temp, _ := strconv.Atoi(string(v))
			num = num*10 + temp
		} else if v == '[' {
			nums, num = append(nums, num), 0
			strs, cur = append(strs, cur), ""
		} else if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			cur += string(v)
		} else if v == ']' {
			for i := 0; i < nums[len(nums)-1]; i++ {
				strs[len(strs)-1] += cur
			}
			cur = strs[len(strs)-1]
			nums = nums[:len(nums)-1]
			strs = strs[:len(strs)-1]
		}
	}
	res += cur
	return res
}
