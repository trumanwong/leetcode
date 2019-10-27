package defangIPaddr

func DefangIPaddr(address string) string {
	addr, ret := []byte(address), []byte{}
	temp := []byte(string("[.]"))
	for _, v := range addr {
		if v == '.' {
			ret = append(ret, temp...)
		} else {
			ret = append(ret, v)
		}
	}
	return string(ret)
}