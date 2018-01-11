package learn_sort

import (
	"sort"
	"testing"
)

func Test_SelectSort(t *testing.T) {
	arr := []int{22, 2, 4, 4, 57, 9, 7234, 523, 23, 7}
	arr2 := SelectSort(arr)
	if !sort.IntsAreSorted(arr2) {
		t.Fatal("排序不正确!!!")
	}
	t.Log("排序正确")
}
