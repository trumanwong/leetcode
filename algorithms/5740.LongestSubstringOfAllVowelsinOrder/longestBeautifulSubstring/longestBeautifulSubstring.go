package longestBeautifulSubstring

import (
	"math"
)

func LongestBeautifulSubstring(word string) int {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	res, length, index := 0, 0, 0
	isStart, isEnd := false, false
	arr := []byte(word)
	for i := 0; i < len(arr); i++ {
		if arr[i] == vowels[0] && !isStart {
			isStart = true
			length = 0
		}
		if isStart {
			if arr[i] != vowels[index] {
				index++
				if index == len(vowels) || arr[i] != vowels[index] {
					if isEnd && index == len(vowels) {
						res = int(math.Max(float64(res), float64(length)))
					}
					isStart, isEnd = false, false
					index, length = 0, 0
					if arr[i] == vowels[0] {
						isStart = true
						length++
					}
				} else {
					length++
				}
			} else {
				length++
			}
			if arr[i] == vowels[4] && !isEnd {
				isEnd = true
			}
		}
	}
	if isStart && isEnd && index >= len(vowels) - 1 {
		res = int(math.Max(float64(res), float64(length)))
	}
	return res
}
