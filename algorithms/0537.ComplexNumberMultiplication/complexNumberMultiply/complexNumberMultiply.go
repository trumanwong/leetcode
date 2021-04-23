package complexNumberMultiply

import (
	"strconv"
	"strings"
)

func ComplexNumberMultiply(a string, b string) string {
	aArr := strings.Split(a, "+")
	bArr := strings.Split(b, "+")

	x1, _ := strconv.Atoi(aArr[0])
	y1, _ := strconv.Atoi(aArr[1][0 : len(aArr[1])-1])
	x2, _ := strconv.Atoi(bArr[0])
	y2, _ := strconv.Atoi(bArr[1][0 : len(bArr[1])-1])
	return strconv.Itoa(x1*x2-y1*y2) + "+" + strconv.Itoa(x1*y2+x2*y1) + "i"
}
