package minAvailableDuration

import "sort"

//你是一名行政助理，手里有两位客户的空闲时间表：slots1 和 slots2，以及会议的预计持续时间 duration，请你为他们安排合适的会议时间。
//
//「会议时间」是两位客户都有空参加，并且持续时间能够满足预计时间 duration 的 最早的时间间隔。
//
//如果没有满足要求的会议时间，就请返回一个 空数组。
//
//
//
//「空闲时间」的格式是 [start, end]，由开始时间 start 和结束时间 end 组成，表示从 start 开始，到 end 结束。
//
//题目保证数据有效：同一个人的空闲时间不会出现交叠的情况，也就是说，对于同一个人的两个空闲时间 [start1, end1] 和 [start2, end2]，要么 start1 > end2，要么 start2 > end1。

func MinAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	m1, m2 := make(map[int]int), make(map[int]int)
	arr1, arr2 := make([]int, len(slots1)), make([]int, len(slots2))
	for i, v := range slots1 {
		m1[v[0]] = v[1]
		arr1[i] = v[0]
	}
	for i, v := range slots2 {
		m2[v[0]] = v[1]
		arr2[i] = v[0]
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	i, j := 0, 0
	res := make([][]int, 0)
	for i < len(arr1) && j < len(arr2) {
		if m1[arr1[i]] < arr2[j] {
			i++
			continue
		}

		if m2[arr2[j]] < arr1[i] {
			j++
			continue
		}

		// 说明存在共同的区间
		large := false
		start, end := arr1[i], m1[arr1[i]]
		if start < arr2[j] {
			start = arr2[j]
		}
		if end > m2[arr2[j]] {
			end = m2[arr2[j]]
			large = true
		}
		if start + duration <= end {
			res = append(res, []int{start, start + duration})
			break
		}
		if !large {
			i++
		}
		j++
	}
	i, j = 0, 0
	for i < len(arr1) && j < len(arr2) {
		if m2[arr2[i]] < arr1[j] {
			i++
			continue
		}

		if m1[arr1[j]] < arr2[i] {
			j++
			continue
		}

		// 说明存在共同的区间
		large := false
		start, end := arr2[i], m2[arr2[i]]
		if start < arr1[j] {
			start = arr1[j]
		}
		if end > m1[arr1[j]] {
			end = m1[arr1[j]]
			large = true
		}
		if start + duration <= end {
			res = append(res, []int{start, start + duration})
			break
		}
		if !large {
			i++
		}
		j++
	}
	if len(res) == 0 {
		return []int{}
	}
	if res[0][0] < res[1][0] {
		return res[0]
	}
	return res[1]
}