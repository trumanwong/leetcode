package bitwiseComplement

import "math"

func BitwiseComplement(N int) int {
	i, M := 1, N
	for M >= 2 {
		i++
		M /= 2
	}
	return int(math.Pow(2, float64(i))) - N - 1
}
