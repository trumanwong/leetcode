package jump

func Jump(nums []int) int {
	lo, hi, steps := 0, 0, 0
	// 查看是否到头
	for ; hi < len(nums) - 1; {
		right := 0
		// 从区间找到最大值
		for i := lo; i <= hi; i++ {
			if i + nums[i] > right {
				right = i + nums[i]
			}
		}
		// 更新区间
		lo = hi + 1
		hi = right
		// 更新步数
		steps++
	}
	return steps
}