package transformArray

import "reflect"

func TransformArray(arr []int) []int {
	temp, t := make([]int, len(arr)), make([]int, len(arr))
	copy(temp, arr)
	for {
		copy(t, temp)
		for i := 1; i < len(arr) - 1; i++ {
			if temp[i] > arr[i - 1] && temp[i] > arr[i + 1] {
				temp[i]--
			}

			if temp[i] < arr[i - 1] && temp[i] < arr[i + 1] {
				temp[i]++
			}
		}
		if reflect.DeepEqual(temp, t) {
			break
		}
		copy(arr, temp)
	}
	return arr
}