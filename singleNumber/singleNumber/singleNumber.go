package singleNumber

func SingleNumber(nums []int) int {
	ret := 0;
	for _, num := range nums {
		ret = ret ^ num
	}
	return ret
}