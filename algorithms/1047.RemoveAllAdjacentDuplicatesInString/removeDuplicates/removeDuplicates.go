package removeDuplicates

func RemoveDuplicates(S string) string {
	if len(S) == 1 {
		return S
	}

	i := 1
	old_str, new_str := S, ""
	for old_str != new_str {
		if i >= len(old_str) {
			i = 1
			old_str, new_str = new_str, ""
		}
		if len(old_str) <= 1 {
			break
		}
		if old_str[i] != old_str[i-1] {
			new_str += string(old_str[i-1])
			if i == len(old_str)-1 {
				new_str += string(old_str[i])
			}
			i++
		} else {
			if i == len(old_str)-2 {
				new_str += string(old_str[i+1])
			}
			i += 2
		}
	}
	return old_str
}
