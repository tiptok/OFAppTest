package sorting

import (
	"log"
	"testing"
)

var input [][]int = [][]int{
	[]int{3, 2, 5, 4, 1},
	[]int{3, 5, 2, 4, 9},
	[]int{4, 6, 8, 5, 9, 7, 11},
	[]int{4, 6, 8, 0, 0, 8, 11, 20, 15},
}

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

func TestBubbleSort(t *testing.T) {
	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		RandomQuickSort(input[i], 0, len(input[i])-1)
		t.Logf("%v out:%v", i, input[i])
	}
}

//插入排序
func TestCountingSort(t *testing.T) {
	var input [][]int = [][]int{
		[]int{3, 2, 5, 4, 1},
		[]int{3, 5, 2, 4, 9},
		[]int{4, 6, 8, 5, 9, 7, 11},
		[]int{4, 6, 8, 0, 0, 8, 11, 20, 15},
	}

	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		CountingSort(input[i])
		t.Logf("%v out:%v", i, input[i])
	}
}

//插入排序
func TestRadixSort(t *testing.T) {
	var input [][]int = [][]int{
		[]int{3, 2, 5, 4, 1},
		[]int{3, 5, 2, 4, 9},
		[]int{4, 6, 8, 5, 9, 7, 11},
		[]int{4, 6, 8, 0, 0, 8, 11, 20, 15},
		[]int{3221, 1, 10, 9680, 577, 9420, 7, 5622, 4793, 2030, 3138, 82, 2599, 743, 4127},
	}

	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		RadixSort(input[i])
		t.Logf("%v out:%v", i, input[i])
	}
}

func TestSwitchCase(t *testing.T) {
	value := 1
	switch value {
	case 1:
		fallthrough
	case 2:
		log.Println("1,2")
		break
	case 3:
		log.Println("3")
		break
	default:
		log.Println("default")
		break
	}
}
