package quickSort

// 分割函数
func division(arr []int, left int, right int) int {
	// 取左侧元素为基准元素
	base := arr[left]
	for left < right {
		// 从右向左查找第一个比基准小的元素
		for left < right && arr[right] > base {
			right--
		}

		// 将比基准小的元素移到左侧
		arr[left] = arr[right]
		// 从左向右查找第一个比基准大的元素
		for left < right && arr[left] < base {
			left++
		}
		// 将比基准大的数移到有段
		arr[right] = arr[left]
	}
	// 保存基准数
	arr[left] = base
	return left
}

// 快速排序
func QuickSort(arr []int, left int, right int) {
	if left < right {
		i := division(arr, left, right)
		// 将基准左侧部分进行递归排序
		QuickSort(arr, left, i - 1)
		// 将基准右侧部分进行排序
		QuickSort(arr, i + 1, right)
	}
}