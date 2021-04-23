package addToArrayForm

func AddToArrayForm(A []int, K int) []int {
	length := len(A)
	cur := K
	res := []int{}
	for i := length - 1; i >= 0 || cur > 0; i-- {
		if i >= 0 {
			cur += A[i]
		}
		res = append(res, cur%10)
		cur = cur / 10
	}
	len_res := len(res)
	loop := len_res / 2
	for i := 0; i < loop; i++ {
		res[i], res[len_res-i-1] = res[len_res-i-1], res[i]
	}
	return res
}
