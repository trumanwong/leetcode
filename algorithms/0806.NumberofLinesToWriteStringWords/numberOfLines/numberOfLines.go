package numberOfLines

func NumberOfLines(widths []int, S string) []int {
	sum, width := 0, 0
	for i := 0; i < len(S); i++ {
		temp := int(S[i] - 'a')
		if sum+widths[temp] <= 100 {
			sum += widths[temp]
		} else {
			width++
			sum = 0
			i--
		}
	}
	if sum != 0 {
		width++
	}
	return []int{width, sum}
}
