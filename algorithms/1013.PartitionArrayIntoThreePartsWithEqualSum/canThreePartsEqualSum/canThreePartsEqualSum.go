package canThreePartsEqualSum

func CanThreePartsEqualSum(A []int) bool {
	sum, count, temp := 0, 0, 0
	for _, v := range A {
		sum += v
	}
	if sum % 3 != 0 {
		return false
	}
	avg := sum / 3
	for _, v := range A {
		temp += v
		if temp == avg {
			count++
			temp = 0
		}
	}
	return count == 3
}