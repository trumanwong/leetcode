package nextGreatestLetter

func NextGreatestLetter(letters []byte, target byte) byte {
	i, j := 0, len(letters)
	if target >= letters[j-1] {
		return letters[0]
	}
	for i < j {
		mid := (j-i)/2 + i
		if letters[mid] <= target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return letters[i]
}
