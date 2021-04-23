package duplicateZeros

func DuplicateZeros(arr []int) {
	dst, index := make([]int, len(arr)), 0
	copy(dst, arr)
	for i := 0; i < len(arr); i++ {
		if dst[index] == 0 {
			arr[i] = 0
			if i+1 < len(arr) {
				arr[i+1] = 0
				i++
			}
		} else {
			arr[i] = dst[index]
		}
		index++
	}
}
