package sorting

import "testing"

//测试堆排序
func TestHeapSort(t *testing.T) {
	var input [][]int = [][]int{
		[]int{4, 6, 8, 5, 9},
		[]int{3, 5, 2, 4, 9},
		[]int{4, 6, 8, 5, 9, 7, 11},
		[]int{4, 6, 8, 0, 0, 7, 11, 20, 15},
	}

	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		HeapSort(input[i])
		t.Logf("%v out:%v", i, input[i])
	}
}

//合并排序
func TestMergeSort(t *testing.T) {
	var input [][]int = [][]int{
		[]int{3, 2, 5, 4, 1},
		[]int{3, 5, 2, 4, 9},
		[]int{4, 6, 8, 5, 9, 7, 11},
		[]int{4, 6, 8, 0, 0, 8, 11, 20, 15},
	}

	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		MergeSort(input[i])
		t.Logf("%v out:%v", i, input[i])
	}
}

//插入排序
func TestInsertionSort(t *testing.T) {
	var input [][]int = [][]int{
		[]int{3, 2, 5, 4, 1},
		[]int{3, 5, 2, 4, 9},
		[]int{4, 6, 8, 5, 9, 7, 11},
		[]int{4, 6, 8, 0, 0, 8, 11, 20, 15},
	}

	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		InsertionSort(input[i])
		t.Logf("%v out:%v", i, input[i])
	}
}

func TestCopy(t *testing.T) {
	a := []int{1}
	b := []int{2}
	copy(a[0:1], b[0:1])
	t.Log(a)
}
