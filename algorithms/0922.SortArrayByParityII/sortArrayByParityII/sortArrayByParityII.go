package sortArrayByParityII

func SortArrayByParityII(A []int) []int {
	j := 1
	for i := 0; i < len(A)-1; i = i + 2 {
		if A[i]&1 != 0 {
			for A[j]&1 != 0 {
				j = j + 2
			}
			tmp := A[i]
			A[i] = A[j]
			A[j] = tmp
		}
	}
	return A
}
