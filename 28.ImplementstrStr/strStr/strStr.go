package strStr

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	i, j := 0, 0
	hsLength := len(haystack)
	needleLength := len(needle)
	for ; i < hsLength && j < needleLength; {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
				// 不匹配， i后退
				i = i - j + 1
				j = 0
		}
	}

	if j == needleLength {
		return i - j
	} else {
		return -1
	}
}