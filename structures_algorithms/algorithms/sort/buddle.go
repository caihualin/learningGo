package learn_sort

// 冒泡排序

func BuddleSort(arr []int) []int {
	l := len(arr) - 1
	for i := 0; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
