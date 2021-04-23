package insertSort

func InsertSort(arr []int) {
	var j int
	for i := 1; i < len(arr); i++ {
		// 取出一个未排序的数据
		temp := arr[i]
		for j = i - 1; j >= 0 && temp < arr[j]; j-- {
			// 向后移动数据
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}
