package countAndSay

import "strconv"

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	cur := "1"
	temp := ""
	var length, idx, count int

	for i := 2; i <= n; i++ {
		length = len(cur)
		count = 1
		temp = ""
		for idx = 1; idx < length; idx++ {
			if length >= 2 && cur[idx] == cur[idx - 1] {
				count++
			} else {
					temp += strconv.Itoa(count)
					temp += string(cur[idx - 1])
					count = 1
			}
		}
		temp += strconv.Itoa(count)
		temp = temp + string(cur[length - 1])
		cur = temp
	}
	return cur
}