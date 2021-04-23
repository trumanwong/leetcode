package findDuplicate

func FindDuplicate(nums []int) int {
	slow, fast, t := 0, 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	for {
		slow = nums[slow]
		t = nums[t]
		if slow == t {
			break
		}
	}
	return slow
}
