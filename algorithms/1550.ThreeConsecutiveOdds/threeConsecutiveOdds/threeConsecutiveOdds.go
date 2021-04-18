package threeConsecutiveOdds

func ThreeConsecutiveOdds(arr []int) bool {
	i := 0
	for _, v := range arr {
		if v % 2 != 0 {
			i++
		} else {
			i = 0
		}
		if i == 3 {
			return true
		}
	}
	return false
}