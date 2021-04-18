package mergeSort

func MergeSort(arr []int, left, right int)  {
	if left >= right {
		return
	}

	mid := left + (right - left) / 2

	MergeSort(arr, left, mid)
	MergeSort(arr, mid + 1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int)  {
	nums := make([]int, right - left + 1)
	p, q := left, mid + 1

	index := 0
	// 比较左右两部分元素，哪个小就将哪个元素复制到nums中
	for p <= mid && q <= right {
		if arr[p] > arr[q] {
			nums[index] = arr[q]
			q++
		} else {
			nums[index] = arr[p]
			p++
		}
		index++
	}
	// 将剩余蒜素依次填入到nums中
	for p <= mid {
		nums[index] = arr[p]
		p++
		index++
	}
	for q <= right {
		nums[index] = arr[q]
		q++
		index++
	}
	// 将最终排序结果复制给原数组
	for i := 0; i < len(nums); i++ {
		arr[left + i] = nums[i]
	}
}