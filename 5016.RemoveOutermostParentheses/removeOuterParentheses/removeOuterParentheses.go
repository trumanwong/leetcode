package removeOuterParentheses

func RemoveOuterParentheses(S string) string {
	res, left, innerLeft, innerRight := "", false, 0, 0
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "(" {
			innerLeft++
		} else {
			innerRight++
		}
		if i + 1 < len(S) && S[i] ^ S[i + 1] == 0 {
			if !left && string(S[i]) == "(" {
				left = true
			} else if left && string(S[i]) == ")" {
				res += string(S[i])
				if innerLeft == innerRight + 1 {
					left, innerRight, innerLeft = false, 0, 0
					i++
				}
			} else {
				res += string(S[i])
			}
		} else if left {
			res += string(S[i])
		}
	}
	return res
}