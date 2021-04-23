package isMajorityElement

func IsMajorityElement(nums []int, target int) bool {
	count := 0
	for _, v := range nums {
		if v == target {
			count++
		} else if v > target {
			break
		}
	}
	if count > len(nums)/2 {
		return true
	}
	return false
}
