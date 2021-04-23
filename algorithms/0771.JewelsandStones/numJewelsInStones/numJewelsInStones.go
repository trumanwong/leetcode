package numJewelsInStones

import "strings"

func NumJewelsInStones(J string, S string) int {
	sum := 0
	for _, j := range J {
		if len(J) == 0 {
			break
		}
		temp := ""
		for _, s := range S {
			if j == s {
				sum++
				temp = string(j)
			}
		}
		if temp != "" {
			S = strings.Replace(S, temp, "", -1)
		}
	}
	return sum
}
