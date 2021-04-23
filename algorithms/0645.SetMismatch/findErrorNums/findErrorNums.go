package findErrorNums

func FindErrorNums(nums []int) []int {
	res, rept, temp := []int{}, 0, make([]int, len(nums))
	for _, v := range nums {
		if temp[v-1] != 0 {
			rept = v
		}
		temp[v-1] = v
	}
	for i, v := range temp {
		if v == 0 {
			res = append(res, []int{rept, i + 1}...)
			break
		}
	}
	return res
}
