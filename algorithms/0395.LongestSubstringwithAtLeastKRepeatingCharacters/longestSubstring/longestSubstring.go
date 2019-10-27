package longestSubstring

import "math"

func LongestSubstring(s string, k int) int {
	if k <= 1 {
		return len(s)
	}

	if len(s) == 0 || k > len(s) {
		return 0
	}

	hash := make([]int, 26)
	for _, v := range s {
		hash[int(v - 'a')]++
	}

	i := 0
	for i < len(s) && hash[int(s[i] - 'a')] >= k {
		i++
	}

	if i == len(s) {
		return len(s)
	}

	l := LongestSubstring(s[0:i], k)
	for i < len(s) && hash[int(s[i] - 'a')] < k {
		i++
	}
	r := LongestSubstring(s[i:], k)
	return int(math.Max(float64(l), float64(r)))
}