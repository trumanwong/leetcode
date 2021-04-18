package count_sort

func CountSort(arr []int) []int {
	res := make([]int, len(arr))
	min, max := arr[0], arr[0]
	for _, v:= range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	// 取排序数组中，元素大小的极值差+1
	k := max - min + 1
	counts := make([]int, k)
	for _, v := range arr {
		counts[v - min + 1]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i - 1]
	}
	for i := len(arr) - 1
}