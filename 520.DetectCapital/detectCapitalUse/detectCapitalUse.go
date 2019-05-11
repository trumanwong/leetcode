package detectCapitalUse

func DetectCapitalUse(word string) bool {
	isFirst, hasLower, hasUpper := false, false, false
	for i, v := range word {
		if i == 0 {
			if v >= 65 && v <= 90 {
				isFirst = true
			}
		} else {
			if v >= 97 && v <= 122 {
				hasLower = true
			} else {
				hasUpper = true
			}
		}
	}
	if isFirst {
		if hasLower && hasUpper {
			return false
		}
	} else if hasUpper {
		return false
	}
	return true
}