package learn_sort

import (
	"sort"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	arr := []int{22, 2, 4, 4, 57, 7234, 1, 523, 23, 7}
	arr2 := QuickSort(arr, 0, len(arr)-1)
	if !sort.IntsAreSorted(arr2) {
		t.Fatal("排序不正确!!!")
	}
	t.Log("排序正确")
}
