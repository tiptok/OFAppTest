package sorting

import "math/rand"

//冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

/*选择排序*/
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			swap(arr, i, min)
		}
	}
}

/*快速排序*/
func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	idx := left - 1
	for i := left; i < right; i++ {
		if arr[i] < pivot {
			idx++
			swap(arr, idx, i)
		}
	}
	swap(arr, idx+1, right)
	return idx + 1
}

/*随机快速排序*/
func RandomQuickSort(arr []int, left, right int) {
	if left < right {
		pivot := randomPartition(arr, left, right)
		RandomQuickSort(arr, left, pivot-1)
		RandomQuickSort(arr, pivot+1, right)
	}
}

func randomPartition(arr []int, left, right int) int {
	r := rand.Intn(right-left) + left
	swap(arr, right, r)
	return partition(arr, left, right)
}
