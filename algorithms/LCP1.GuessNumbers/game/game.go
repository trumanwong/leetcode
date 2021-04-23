package game

func game(guess []int, answer []int) int {
	ret := 0
	for i, v := range guess {
		if v == answer[i] {
			ret++
		}
	}
	return ret
}
