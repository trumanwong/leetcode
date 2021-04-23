package longestPalindrome

func LongestPalindrome(s string) int {
	count := make([]int, 128)
	for _, v := range s {
		count[int(v)]++
	}

	res := 0
	for _, v := range count {
		res += v / 2 * 2
		if v%2 == 1 && res%2 == 0 {
			res++
		}
	}
	return res
}
