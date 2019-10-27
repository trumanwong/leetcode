package intToRoman

//Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
//Symbol       Value
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
//
//Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
//I can be placed before V (5) and X (10) to make 4 and 9.
//X can be placed before L (50) and C (100) to make 40 and 90.
//C can be placed before D (500) and M (1000) to make 400 and 900.
//Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
//
//Example 1:
//
//Input: 3
//Output: "III"
//Example 2:
//
//Input: 4
//Output: "IV"
//Example 3:
//
//Input: 9
//Output: "IX"
//Example 4:
//
//Input: 58
//Output: "LVIII"
//Explanation: L = 50, V = 5, III = 3.
//Example 5:
//
//Input: 1994
//Output: "MCMXCIV"
//Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
func IntToRoman(num int) string {
	maps := map[int]string {
		1: "I",
		5: "V",
		10: "X",
		50: "L",
		100: "C",
		500: "D",
		1000: "M",
	}
	result := ""
	for i, v := range maps {
		if i == num {
			return v
		}
	}
	if num == 4 {
		return "IV"
	} else if num == 9 {
		return "IX"
	} else if num == 40 {
		return "XL"
	} else if num == 90 {
		return "XC"
	} else if num == 400 {
		return "CD"
	} else if num == 900 {
		return "CM"
	} else if num < 4 {
		for i := 0; i < num; i++ {
			result = result + maps[1]
		}
	} else if num < 10 {
		result += maps[5]
		for i := 0; i < num - 5; i++ {
			result += maps[1]
		}
	} else if num < 50 {
		mod := num % 10
		if num > 40 {
			result += "XL"
			mod = num - 40
		} else {
			for i := 0; i < num/10; i++ {
				result += maps[10]
			}
		}
		result += IntToRoman(mod)
	} else if num < 100 {
		sub := num - 50
		if num > 90 {
			sub = num - 90
			result += "XC"
		} else {
			result += maps[50]
		}
		result += IntToRoman(sub)
	} else if num < 500 {
		mod := num % 100
		if num > 400 {
			result += "CD"
			mod = num - 400
		} else {
			for i := 0; i < num / 100; i++ {
				result += maps[100]
			}
		}
		result += IntToRoman(mod)
	} else if num < 1000 {
		sub := num - 500
		if num > 900 {
			result += "CM"
			sub = num - 900
		} else {
			result += maps[500]
		}
		result += IntToRoman(sub)
	} else {
		for i := 0; i < num / 1000; i++ {
			result += maps[1000]
		}
		mod := num % 1000
		result += IntToRoman(mod)
	}
	return result
}
