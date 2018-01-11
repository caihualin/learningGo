package learn_search

import (
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if BinarySearch(arr, 7) {
		t.Log("存在9")
	}
	if BinarySearch(arr, 100) {
		t.Log("存在100")
	}
}
