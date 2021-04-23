package gcdOfStrings

import "strings"

func GcdOfStrings(str1 string, str2 string) string {
	S, T := str1, str2
	if len(str2) > len(str1) {
		S, T = str2, str1
	}
	b1, b2 := []byte(S), []byte(T)
	for i, v := range b2 {
		if v != b1[i] {
			return ""
		}
	}

	if len(S) == len(T) {
		return S
	}

	res, temp, gcdVal := "", []byte{}, gcd(len(S), len(T))
	for i, v := range b2 {
		if i >= gcdVal {
			break
		}
		temp = append(temp, v)
		if len(S)%len(temp) == 0 && len(T)%len(temp) == 0 {
			if strings.Repeat(string(temp), len(S)/len(temp)) == S && strings.Repeat(string(temp), len(T)/len(temp)) == T {
				res = string(temp)
			}
		}
	}
	return res
}

func gcd(i, j int) int {
	for i != j {
		if i > j {
			i -= j
		} else {
			j -= i
		}
	}
	return i
}
