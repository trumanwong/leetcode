package wordBreak

import "strings"

func WordBreak(s string, wordDict []string) bool {
	dp := make([]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(wordDict); j++ {
			if strings.Index(s[i:], wordDict[j]) == 0 {
				k := i + len(wordDict[j])
				if k == len(s) || dp[k] == 1 {
					dp[i] = 1
					break
				}
			}
		}
	}
	return dp[0] == 1
}
