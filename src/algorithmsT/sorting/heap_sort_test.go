package sorting

import "testing"

func TestHeapSort(t *testing.T) {
	var input [][]int = [][]int{
		[]int{4, 6, 8, 5, 9},
		[]int{4, 6, 8, 5, 9, 7, 11},
	}

	for i := range input {
		t.Logf("%v in:%v", i, input[i])
		HeapSort(input[i])
		t.Logf("%v out:%v", i, input[i])
	}
}
