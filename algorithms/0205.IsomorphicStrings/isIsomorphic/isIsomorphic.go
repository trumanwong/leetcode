package isIsomorphic

func IsIsomorphic(s string, t string) bool {
	ms, mt, sByte := make(map[byte]byte), make(map[byte]byte), []byte(s)
	for i, v := range sByte {
		if _, ok := ms[v]; ok{
			if ms[v] != t[i] {
				return false
			}
		} else {
			ms[v] = t[i]
			mt[t[i]]++
			if mt[t[i]] > 1 {
				return false
			}
		}
	}
	return true
}