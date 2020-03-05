package sorting

//堆排序
// 0 1 2 3 4 5 6
//[4 6 8 5 9 7 11]
func HeapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
}

func adjustHeap(arr []int, i, len int) {
	tmp := arr[i]
	for j := i*2 + 1; i < len; j = j*2 + 1 {
		if (j+1) < len && arr[j] < arr[j+1] {
			j++
		}
		if arr[j] > tmp {
			arr[i] = arr[j]
			i = j
		} else {
			break
		}
	}
	arr[i] = tmp
}
