package singleNumberIII

func SingleNumber(nums []int) []int {
	diff := 0
	for _, v := range nums {
		diff = diff ^ v
	}
	//得到的diff表示两个不同的数的按位异或的结果
	diff = -diff & diff
	result := make([]int, 2)
	for _, v := range nums {
		if diff & v == 0 {
			result[0] = result[0] ^ v
		} else {
			result[1] = result[1] ^ v
		}
	}
	return result
}