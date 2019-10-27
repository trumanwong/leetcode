package canJump

func CanJump(nums []int) bool {
	far := 0
	for i, v := range nums {
		if far < i + v {
			far = i + v
		}
		if far <= i && i != len(nums) - 1 {
			return false
		}
	}
	return true
}