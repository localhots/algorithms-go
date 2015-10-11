package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testSort(t, SelectionSort,
		[]int{5, 3, 1, 4, 2},
		[]int{1, 2, 3, 4, 5},
	)
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(unsortedArray())
	}
}
