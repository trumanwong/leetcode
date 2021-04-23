package threeSum

import "sort"

func ThreeSum(nums []int) [][]int {
	/*
	   排序;
	   确定a，寻找 a + b + c = 0;
	   双指针即为 b, c;
	   注意手动去重；因为确定三数和为0，那么a>0即可提前结束；
	*/
	res := make([][]int, 0)
	sort.Ints(nums)

	for i, v := range nums {
		if v > 0 {
			break
		}

		if i > 0 && v == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if r < len(nums)-1 && nums[r] == nums[r+1] || v+nums[l]+nums[r] > 0 {
				r--
			} else if l > i+1 && nums[l] == nums[l-1] || v+nums[l]+nums[r] < 0 {
				l++
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return res
}
