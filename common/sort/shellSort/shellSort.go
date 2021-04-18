package shellSort

func ShellSort(arr[] int)  {
	// 计算第1次的增量
	d := len(arr) / 2
	// 循环至增量为1时结束
	for d >= 1 {
		for i := d; i < len(arr); i++ {
			// 获取序列中的下一个数据
			next := arr[i]
			// 序列中前一个数据的序号
			j := i - d
			// 下一个数大于钱一个属
			for j >= 0 && arr[j] > next {
				arr[j + d] = arr[j]
				j = j - d
			}
			// 保存数据
			arr[j + d] = next
		}
		// 缩小增量
		d /= 2
	}
}