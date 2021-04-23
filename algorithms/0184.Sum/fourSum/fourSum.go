package fourSum

import "sort"

func FourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}

	sort.Ints(nums)
	length, threeEnd, twoEnd := len(nums), len(nums)-3, len(nums)-2
	for i := 0; i < threeEnd; i++ {
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			return res
		}

		if (i > 0 && nums[i-1] == nums[i]) || nums[i]+nums[length-3]+nums[length-2]+nums[length-1] < target {
			continue
		}

		for j := i + 1; j < twoEnd; j++ {
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if (j > i+1 && nums[j-1] == nums[j]) || (nums[i]+nums[j]+nums[length-2]+nums[length-1] < target) {
				continue
			}

			l, r := j+1, length-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for nums[l] == nums[l+1] && l+1 < r {
						l++
					}

					for nums[r] == nums[r-1] && l < r-1 {
						r--
					}
					l++
					r--
				}
			}
		}
	}
	return res
}
