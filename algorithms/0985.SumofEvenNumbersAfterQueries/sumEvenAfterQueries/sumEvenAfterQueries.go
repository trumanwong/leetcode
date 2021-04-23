package sumEvenAfterQueries

func SumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}
	res := make([]int, len(queries))
	for i, v := range queries {
		val, index := v[0], v[1]
		if A[index]%2 == 0 {
			sum -= A[index]
		}
		A[index] += val
		if A[index]%2 == 0 {
			sum += A[index]
		}
		res[i] = sum
	}
	return res
}
