package minOperations

func MinOperations(logs []string) int {
	res := make([]string, 0)
	for _, v := range logs {
		if v == "./" {
			continue
		}
		if v == "../" {
			if len(res) > 0 {
				res = res[0 : len(res)-1]
			}
		} else {
			res = append(res, v)
		}
	}
	return len(res)
}
