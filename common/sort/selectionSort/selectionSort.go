package selectionSort

func SelectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		// 记录需要和第k个元素交换
		k := i
		for j := i + 1; j < n; j++ {
			// 查找最小数据，保存其序号在k中
			if arr[k] > arr[j] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
}
