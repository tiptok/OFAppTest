package sorting

//堆排序
// 0 1 2 3 4 5 6
//[4 6 8 5 9 7 11]
func HeapSort(arr []int) {
	//1.调整最大堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeapNode(arr, i, len(arr))
	}

	//2.交换堆顶元素与末尾元素
	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		adjustHeapNode(arr, 0, i)
	}
}

//大堆正序  小堆倒序
func adjustHeapNode(arr []int, i, len int) {
	tmp := arr[i]
	for j := i*2 + 1; j < len; j = j*2 + 1 {
		//if (j+1) < len && arr[j] < arr[j+1] {
		//	j++
		//}
		//if arr[j] > tmp {  //arr[j]>tmp 大堆   arr[j]<tmp 小堆
		//	arr[i] = arr[j]
		//	i = j
		//} else {
		//	break
		//}

		if (j+1) < len && arr[j] > arr[j+1] {
			j++
		}
		if arr[j] < tmp { //arr[j]>tmp 大堆   arr[j]<tmp 小堆
			arr[i] = arr[j]
			i = j
		} else {
			break
		}
	}
	arr[i] = tmp
	//log.Println(arr)
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
