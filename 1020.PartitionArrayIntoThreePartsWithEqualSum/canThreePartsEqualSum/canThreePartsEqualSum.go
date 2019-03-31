package canThreePartsEqualSum

func CanThreePartsEqualSum(A []int) bool {
	res := false
	index := 2
	ALength := len(A)
	for !res {
		if index >= ALength {
			break
		}
		temp := 0
		for i := 0; i < index; i++ {
			temp += A[i]
		}
		index2 := 0
		sum := 0
		if temp == A[index] {
			index++
		}
		for i := index; i < ALength; i++ {
			sum += A[i]
			if sum == temp {
				index2 = i + 1
				break
			}
		}
		if index2 == 0 || index2 == ALength {
			index++
			continue
		}
		sum = 0
		if temp == index2 {
			index2++
		}
		for i := index2; i < ALength; i++ {
			sum += A[i]
			if sum == temp {
				res = true
				break
			}
		}
		index++
	}
	return res
}