package sort

import (
	"testing"
)

func TestShuffleSort(t *testing.T) {
	testSort(t, ShuffleSort,
		[]int{5, 3, 1, 4, 2},
		[]int{1, 2, 3, 4, 5},
	)
}
