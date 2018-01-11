package learn_sort

// 快速排序

func QuickSort(arr []int, start, end int) []int {
	key := arr[(start+end)/2]
	left, right := start, end
	for left < right {
		for arr[left] < key {
			left++
		}

		for arr[right] > key {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if start < right {
		QuickSort(arr, start, right)
	}
	if end > left {
		QuickSort(arr, left, end)
	}

	return arr
}
