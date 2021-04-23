package heapSort

// 构成堆
// arr 数组
// s 需要构成堆的根节点序号
// n 数组长度
func heapAdjust(arr []int, s int, n int) {
	// 第s个节点有右子树
	for 2*s+1 < n {
		j := 2*s + 1
		// 如果左子树小于右子树，则需要比较右子树
		if j+1 < n && arr[j] < arr[j+1] {
			j++
		}
		if arr[s] < arr[j] {
			// 如果父节点小于子节点，则交换父子节点
			arr[s], arr[j] = arr[j], arr[s]
			// 堆被破坏，需要重新调整
			s = j
		} else {
			// 不再需要调整
			break
		}
	}
}

func HeapSort(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapAdjust(arr, i, n)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		// 将arr[0]至arr[i]重新调整为堆
		heapAdjust(arr, 0, i)
	}
}
