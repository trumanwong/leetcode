package compress

import "strconv"

func Compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	count, cur, index := 1, chars[0], 0
	for i := 1; i < len(chars); i++ {
		if chars[i] == cur {
			count++
		} else {
			chars[index] = cur
			index++
			if count > 1 {
				temp := []byte(strconv.Itoa(count))
				for _, v := range temp {
					chars[index] = v
					index++
				}
			}
			cur = chars[i]
			count = 1
		}
	}
	chars[index] = cur
	index++
	if count > 1 {
		temp := []byte(strconv.Itoa(count))
		for _, v := range temp {
			chars[index] = v
			index++
		}
	}
	chars = chars[0:index]
	return len(chars)
}
