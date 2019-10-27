package addBinary

import (
	"math"
)

func AddBinary(a string, b string) string {
	isCarry := false
	res := ""
	zeros := ""
	diff := len(a) - len(b)
	abs_diff := int(math.Abs(float64(diff)))
	for i := 0; i < abs_diff; i++ {
		zeros += "0"
	}
	if diff < 0 {
		a = zeros + a
	} else {
		b = zeros + b
	}
	length := len(a)
	for i := length - 1; i >= 0; i-- {
		tempA := string(a[i])
		tempB := string(b[i])
		add := "0"
		if tempA == "1" && tempB == "1" {
			if isCarry {
				add = "1"
			}
			isCarry = true
		} else if tempA == "1" || tempB == "1" {
			if !isCarry {
				add = "1"
				isCarry = false
			} else {
				add = "0"
				isCarry = true
			}
		} else {
			if isCarry {
				add = "1"
			} else {
				add = "0"
			}
			isCarry = false
		}
		res = add + res
	}
	if isCarry {
		res = "1" + res
	}
	return res
}