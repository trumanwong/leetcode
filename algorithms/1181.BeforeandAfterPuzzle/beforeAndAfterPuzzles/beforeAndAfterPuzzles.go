package beforeAndAfterPuzzles

import (
	"sort"
	"strings"
)

func BeforeAndAfterPuzzles(phrases []string) []string {
	m := make(map[string]int)
	phras := [][]string{}
	for i := 0; i < len(phrases); i++ {
		temp := strings.Split(phrases[i], " ")
		phras = append(phras, temp)
	}

	for i := 0; i < len(phrases); i++ {
		for j := i + 1; j < len(phrases); j++ {
			str1, str2 := "", ""
			if phras[i][len(phras[i])-1] == phras[j][0] {
				str1 = phrases[i][0:(len(phrases[i])-len(phras[i][len(phras[i])-1]))] + phrases[j]
			}

			if phras[j][len(phras[j])-1] == phras[i][0] {
				str2 = phrases[j][0:(len(phrases[j])-len(phras[j][len(phras[j])-1]))] + phrases[i]
			}
			m[str1]++
			m[str2]++
		}
	}

	delete(m, "")
	ret := []string{}
	for k, _ := range m {
		ret = append(ret, k)
	}

	sort.Strings(ret)
	return ret
}
