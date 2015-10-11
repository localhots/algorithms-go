package sort

import (
	"testing"
)

func TestBinaryInsertionSort(t *testing.T) {
	testSort(t, BinaryInsertionSort,
		[]int{5, 3, 1, 4, 2},
		[]int{1, 2, 3, 4, 5},
	)
}

func BenchmarkBinaryInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinaryInsertionSort(unsortedArray())
	}
}
