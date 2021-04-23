package getRow

func GetRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	num := 1

	for i := 0; i <= rowIndex/2; i++ {
		result[i] = num
		result[rowIndex-i] = num
		num = num * (rowIndex - i) / (i + 1)
	}
	return result
}
