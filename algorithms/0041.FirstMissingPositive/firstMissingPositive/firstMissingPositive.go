package firstMissingPositive

func FirstMissingPositive(nums []int) int {
	numSize := len(nums)
	if numSize < 1 {
		return 1
	}

	i := 0
	// 2 1 3 4
	for i < numSize {
		// x大于0，x不等于当前位置索引，x小于数组大小，x不等于n[x]的值
		if nums[i] != i+1 && nums[i] > 0 && nums[i] <= numSize && nums[i] != nums[nums[i]-1] {
			// 交换x和n[x]
			temp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[temp-1] = temp
		} else {
			i++
		}
	}

	// 寻找第一个空位
	for i = 0; i < numSize; i++ {
		if nums[i] != i+1 {
			break
		}
	}
	return i + 1
}
