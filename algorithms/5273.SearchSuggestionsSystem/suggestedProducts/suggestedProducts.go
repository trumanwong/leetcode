package suggestedProducts

import (
	"sort"
	"strings"
)

func SuggestedProducts(products []string, searchWord string) [][]string {
	res := make([][]string, 0)
	sort.Strings(products)

	last := 0
	for i := 0; i < len(searchWord); i++ {
		prefix := searchWord[0:i+1]
		prefixs := make([]string, 0)
		for j := last; j < len(products); j++ {
			if strings.HasPrefix(products[j], prefix) {
				if len(prefix) == 0 {
					last = j
				}
				prefixs = append(prefixs, products[j])
			}
			if len(prefixs) == 3 {
				break
			}
		}
		res = append(res, prefixs)
	}
	return res
}