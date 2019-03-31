package reverseString

func ReverseString(s []byte)  {
	length := len(s)
	loop := len(s) / 2
	for i := 0; i < loop; i++ {
		temp := s[i]
		s[i] = s[length - i - 1]
		s[length - i - 1] = temp
	}
}