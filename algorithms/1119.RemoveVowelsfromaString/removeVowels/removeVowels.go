package removeVowels

func RemoveVowels(S string) string {
	s := []byte(S)
	ret := []byte{}
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			continue
		}
		ret = append(ret, v)
	}
	return string(ret)
}
