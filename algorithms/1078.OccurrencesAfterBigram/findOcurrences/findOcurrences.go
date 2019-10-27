package findOcurrences

func FindOcurrences(text string, first string, second string) []string {
	textBytes, textArr := []byte(text), []string{}
	start := -1
	for i, v := range textBytes {
		if v != ' ' && start == -1 {
			start = i
		}
		if v == ' ' && start != -1 {
			textArr = append(textArr, string(textBytes[start:i]))
			start = -1
		}
	}
	if start != -1 {
		textArr = append(textArr, string(textBytes[start:len(textBytes)]))
	}
	res := []string{}
	for i := 0; i < len(textArr) - 2; i++ {
		if textArr[i] == first && textArr[i + 1] == second {
			res = append(res, textArr[i + 2])
		}
	}
	return res
}