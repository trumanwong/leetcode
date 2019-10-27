package judgePoint24

func JudgePoint24(nums []int) bool {
	new_nums := []float64{}
	for _, v := range nums {
		new_nums = append(new_nums, float64(v))
	}
	return solve(new_nums)
}

func solve(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return myAbs(nums[0] - 24) < 0.0000000001
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				nums2 := []float64{}
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						nums2 = append(nums2, nums[k])
					}
				}

				for k := 0; k < 4; k++ {
					if k < 2 && j > i {
						continue
					}
					if k == 0 {
						nums2 = append(nums2, nums[i] + nums[j])
					} else if k == 1 {
						nums2 = append(nums2, nums[i] * nums[j])
					} else if k == 2 {
						nums2 = append(nums2, nums[i] - nums[j])
					} else {
						if nums[j] != 0 {
							nums2 = append(nums2, nums[i] / nums[j])
						} else {
							continue
						}
					}
					if solve(nums2) {
						return true
					}
					nums2 = nums2[:len(nums2) - 1]
				}
			}
		}
	}
	return false
}

func myAbs(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}