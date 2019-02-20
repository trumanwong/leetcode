/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
package reverse

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	result := 0
	temp := 0
	mul := -1
	if x / mul < 0 {
		mul = 1
	} else {
		x = x / mul
	}
	xLen := len(strconv.Itoa(x))
	for ; xLen > 0; {
		temp = x % 10
		result = result + temp * int(math.Pow(float64(10), float64(xLen - 1)))
		x = x / 10
		if xLen == 1 {
			temp = x / 10
			result = result + temp * int(math.Pow(float64(10), float64(0)))
			xLen--
		}
		xLen--
	}
	result = result * mul
	if result < int(math.Pow(-2, 31)) || result > int(math.Pow(2, 31) - 1) {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(1534236469))
}
