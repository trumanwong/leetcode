package checkIfPangram

func CheckIfPangram(sentence string) bool {
	m := make(map[byte]bool)
	var i byte
	for i = 97; i <= 122; i++ {
		m[i] = false
	}
	bytes := []byte(sentence)
	for j := 0; j < len(bytes); j++ {
		m[bytes[j]] = true
	}
	result := true
	for i = 97; i <= 122; i++ {
		if !m[i] {
			result = false
			break
		}
	}
	return result
}
