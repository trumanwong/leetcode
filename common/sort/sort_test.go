package main

import (
	"leetcode/common/sort/quickSort"
	"leetcode/common/sort/selectionSort"
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