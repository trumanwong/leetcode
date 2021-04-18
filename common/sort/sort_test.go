package main

import (
	"leetcode/common/sort/countSort"
	"leetcode/common/sort/heapSort"
	"leetcode/common/sort/insertSort"
	"leetcode/common/sort/mergeSort"
	"leetcode/common/sort/quickSort"
	"leetcode/common/sort/selectionSort"
	"leetcode/common/sort/shellSort"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T)  {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	origin_sort, origin := append([]int{}, arr...), append([]int{}, arr...)
	sort.Ints(origin_sort)
	quickSort.QuickSort(arr, 0, 9)
	if !reflect.DeepEqual(origin_sort, arr) {
		t.Errorf("Got %v for input %v; expected %v", arr, origin, origin_sort)
	}
}

func TestSelectionSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	origin_sort, origin := append([]int{}, arr...), append([]int{}, arr...)
	sort.Ints(origin_sort)
	selectionSort.SelectionSort(arr, len(arr))
	if !reflect.DeepEqual(origin_sort, arr) {
		t.Errorf("Got %v for input %v; expected %v", arr, origin, origin_sort)
	}
}

func TestHeapSort(t *testing.T)  {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	originSort, origin := append([]int{}, arr...), append([]int{}, arr...)
	sort.Ints(originSort)
	heapSort.HeapSort(arr, len(arr))
	if !reflect.DeepEqual(originSort, arr) {
		t.Errorf("Got %v for input %v; expected %v", arr, origin, originSort)
	}
}

func TestInsertSort(t *testing.T)  {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	originSort, origin := append([]int{}, arr...), append([]int{}, arr...)
	sort.Ints(originSort)
	insertSort.InsertSort(arr)
	if !reflect.DeepEqual(originSort, arr) {
		t.Errorf("Got %v for input %v; expected %v", arr, origin, originSort)
	}
}

func TestShellSort(t *testing.T)  {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	originSort, origin := append([]int{}, arr...), append([]int{}, arr...)
	sort.Ints(originSort)
	shellSort.ShellSort(arr)
	if !reflect.DeepEqual(originSort, arr) {
		t.Errorf("Got %v for input %v; expected %v", arr, origin, originSort)
	}
}

func TestMergeSort(t *testing.T)  {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	originSort, origin := append([]int{}, arr...), append([]int{}, arr...)
	sort.Ints(originSort)
	mergeSort.MergeSort(arr, 0, len(arr) - 1)
	if !reflect.DeepEqual(originSort, arr) {
		t.Errorf("Got %v for input %v; expected %v", arr, origin, originSort)
	}
}

func TestCountSort(t *testing.T)  {
	rand.Seed(time.Now().Unix())
	var arr []int
	for i := 0; i < 10; i++ {
		temp := rand.Intn(1000)
		arr = append(arr, temp)
	}
	originSort := append([]int{}, arr...)
	sort.Ints(originSort)
	res := countSort.CountSort(arr)
	if !reflect.DeepEqual(originSort, res) {
		t.Errorf("Got %v for input %v; expected %v", res, arr, originSort)
	}
}