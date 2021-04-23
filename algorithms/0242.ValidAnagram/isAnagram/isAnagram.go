package isAnagram

func IsAnagram(s string, t string) bool {
	ms, mt := make(map[int]int), make(map[int]int)
	for _, v := range s {
		ms[int(v)]++
	}
	for _, v := range t {
		mt[int(v)]++
	}
	if len(ms) != len(mt) {
		return false
	}
	for i, v := range ms {
		if _, ok := mt[i]; ok && mt[i] == v {
			continue
		} else {
			return false
		}
	}
	return true
}
