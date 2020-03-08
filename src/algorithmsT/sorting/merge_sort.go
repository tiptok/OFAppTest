package sorting

//归并排序法
func MergeSort(arr []int) {
	tmp := make([]int, len(arr))
	mergeSort(arr, 0, len(arr)-1, tmp)
}
func mergeSort(arr []int, left, right int, tmp []int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	//左边排序
	mergeSort(arr, left, mid, tmp)
	//右边排序
	mergeSort(arr, mid+1, right, tmp)
	//合并
	l, r := left, mid+1
	index := 0
	for {
		if !(l <= mid && r <= right) {
			break
		}
		if arr[l] <= arr[r] {
			tmp[index] = arr[l]
			l++
		} else {
			tmp[index] = arr[r]
			r++
		}
		index++
	}

	//合并剩余
	for {
		if !(l <= mid) {
			break
		}
		tmp[index] = arr[l]
		l++
		index++
	}
	for {
		if !(r <= right) {
			break
		}
		tmp[index] = arr[r]
		r++
		index++
	}
	//合并tmp到arr
	index = 0
	for {
		if !(left <= right) {
			break
		}
		arr[left] = tmp[index]
		left++
		index++
	}

}
