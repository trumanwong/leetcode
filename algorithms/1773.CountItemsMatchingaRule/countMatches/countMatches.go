package countMatches

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	rules := map[string]int{"type": 0, "color": 1, "name": 2}
	for _, item := range items {
		if item[rules[ruleKey]] == ruleValue {
			res++
		}
	}
	return res
}
