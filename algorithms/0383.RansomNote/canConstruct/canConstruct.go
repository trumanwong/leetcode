package canConstruct

func CanConstruct(ransomNote string, magazine string) bool {
	res := make([]int, 26)
	for _, v := range magazine {
		res[v-'a']++
	}
	for _, v := range ransomNote {
		res[v-'a']--
		if res[v-'a'] < 0 {
			return false
		}
	}
	return true
}
