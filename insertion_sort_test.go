package algorithms

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testSort(t, InsertionSort,
		[]int{5, 3, 1, 4, 2},
		[]int{1, 2, 3, 4, 5},
	)
}

func TestInsertionSortOptimized(t *testing.T) {
	testSort(t, InsertionSortOptimized,
		[]int{5, 3, 1, 4, 2},
		[]int{1, 2, 3, 4, 5},
	)
}

func TestInsertionSortOptimizedTwoItems(t *testing.T) {
	testSort(t, InsertionSortOptimized,
		[]int{2, 1},
		[]int{1, 2},
	)
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(unsortedArray())
	}
}

func BenchmarkInsertionSortOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSortOptimized(unsortedArray())
	}
}
