package numUniqueEmails

import "strings"

func NumUniqueEmails(emails []string) int {
	res := make(map[string]int)
	for _, v := range emails {
		temps := strings.Split(v, "@")
		name := strings.Replace(strings.Split(temps[0], "+")[0], ".", "", -1)
		res[name+"@"+temps[1]] = 1
	}
	return len(res)
}
