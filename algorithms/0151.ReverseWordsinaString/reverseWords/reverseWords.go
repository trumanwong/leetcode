package reverseWords

func ReverseWords(s string) string {
	S, ret := []byte(s), []byte{}
	bytes := [][]byte{}
	start := -1
	for i, v := range S {
		if v != ' ' && start == -1 {
			start = i
		} else if v == ' ' && start != -1 {
			temp := S[start:i]
			bytes = append(bytes, temp)
			start = -1
		}
	}
	if start != -1 {
		temp := S[start:len(S)]
		bytes = append(bytes, temp)
	}
	for i := len(bytes) - 1; i >= 0; i-- {
		ret = append(ret, bytes[i]...)
		if i != 0 {
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}