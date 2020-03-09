package sorting

//插入排序
func InsertionSort(arr []int) {
	insertionSort(arr, 0, len(arr)-1)
}

func insertionSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	for i := left; i < right; i++ {
		cur := arr[i]  //当前值
		nv := arr[i+1] //下一个值
		if cur < nv {
			continue
		}
		for j := i; j >= 0; j-- {
			if nv < arr[j] {
				swap(arr, j, j+1)
				continue
			}
			break
		}
	}
}
