package singleNonDuplicate

func SingleNonDuplicate(nums []int) int {
	l, r := 0, len(nums) - 1
	for l < r {
		m := l + (r-l)/2
		if m & 1 == 1 { //如果m坐标值为奇数则是偶数个数字
			m--
		}
		if nums[m] != nums[m+1] {
			r = m
		} else {
			l = m + 2
		}
	}
	return nums[l]
}