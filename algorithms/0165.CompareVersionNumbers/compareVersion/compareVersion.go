package compareVersion

import "strconv"

func CompareVersion(version1 string, version2 string) int {
	v1 := split([]byte(version1))
	v2 := split([]byte(version2))
	i := 0
	if len(v1) < len(v2) {
		for j := len(v1); j < len(v2); j++ {
			v1 = append(v1, 0)
		}
	} else {
		for j := len(v2); j < len(v1); j++ {
			v2 = append(v2, 0)
		}
	}
	for i < len(v1) && i < len(v2) {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
		i++
	}
	return 0
}

func split(str []byte) []int {
	start, res := -1, []int{}
	for i, v := range str {
		if v != '.' && start == -1 {
			start = i
		} else if v == '.' && start != -1 {
			temp, _ := strconv.Atoi(string(str[start:i]))
			res = append(res, temp)
			start = -1
		}
	}
	if start != -1 {
		temp, _ := strconv.Atoi(string(str[start:len(str)]))
		res = append(res, temp)
	}
	return res
}
