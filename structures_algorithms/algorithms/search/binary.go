package learn_search

/* 二分查找：
条件：有序元素
*/

func BinarySearch(arr []int, n int) bool {
	var start, end, mid int

	start, end = 0, len(arr)-1
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] == n {
			return true
		}
		if arr[mid] < n {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
