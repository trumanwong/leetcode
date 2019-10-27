package moveZeroes

func MoveZeroes(nums []int)  {
	mov := 0
	for _, v := range nums {
		if v != 0 {
			nums[mov] = v
			mov++
		}
	}
	for mov < len(nums) {
		nums[mov] = 0
		mov++
	}
}