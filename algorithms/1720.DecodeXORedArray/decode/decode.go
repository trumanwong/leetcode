package decode

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i, v := range encoded {
		res[i+1] = res[i] ^ v
	}
	return res
}
