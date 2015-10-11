package sort

import (
	"testing"
)

func TestSleepSort(t *testing.T) {
	testSort(t, SleepSort,
		[]int{5, 3, 1, 4, 2},
		[]int{1, 2, 3, 4, 5},
	)
}
