package findMaxConsecutiveOnes

func FindMaxConsecutiveOnes(nums []int) int {
	countOne, res := 0, 0
	for _, v := range nums {
		if v == 1 {
			countOne++
			if countOne > res {
				res = countOne
			}
		} else {
			countOne = 0
		}
	}
	return res
}