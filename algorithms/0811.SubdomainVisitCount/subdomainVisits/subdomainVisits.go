package subdomainVisits

import "strconv"

func SubdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, value := range cpdomains {
		index, start := -1, -1
		temp := []byte(value)
		count := 0
		for i, v := range temp {
			if v == ' ' {
				index = i
				count, _ = strconv.Atoi(string(value[0:index]))
			} else if index != -1 {
				if start == -1 && v != '.' {
					start = i
				} else if v == '.' && start != -1 {
					domain := string(temp[start:i]) + string(temp[i:len(temp)])
					m[domain] += count
					start = -1
				}
			}
		}

		if start != -1 {
			domain := string(temp[start:len(temp)])
			m[domain] += count
		}
	}

	ret := []string{}
	for k, v := range m {
		count := strconv.Itoa(v)
		ret = append(ret, count+" "+k)
	}
	return ret
}
