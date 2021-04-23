package matrixReshape

func MatrixReshape(nums [][]int, r int, c int) [][]int {
	row, col := len(nums), len(nums[0])
	if row*col != r*c {
		return nums
	}

	res, mul := make([][]int, r), r*c
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i := 0; i < mul; i++ {
		res[i/c][i%c] = nums[i/col][i%col]
	}
	return res
}
