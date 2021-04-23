package findTheDifference

func FindTheDifference(s string, t string) byte {
	var res byte = t[0]
	for i, v := range []byte(s) {
		res ^= v
		res ^= t[i+1]
	}
	return res
}
